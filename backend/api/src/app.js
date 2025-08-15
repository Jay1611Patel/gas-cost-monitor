require('dotenv').config({ path: '../../.env' });

const express = require("express");
const bodyParser = require('body-parser');
const mongoose = require('mongoose');
const { ethers } = require('ethers');
const { SiweMessage } = require('siwe');
const session = require('express-session');
const User = require('../models/User');
const GasReport = require('../models/GasReport');
const { Octokit } = require('@octokit/rest');
const jwt = require('jsonwebtoken');
const fs = require('fs');

const app = express();
const PORT = process.env.API_PORT || 3000;
const MONGO_URI = process.env.MONGO_URI || 'mongodb://localhost:27017/gas_costs';
const JWT_SECRET = 'super-secret-jwt-key';

app.use(express.json());
app.use(bodyParser.json());

const GITHUB_APP_ID = process.env.GITHUB_APP_ID;
const GITHUB_CLIENT_ID = process.env.GITHUB_CLIENT_ID;
const GITHUB_CLIENT_SECRET = process.env.GITHUB_CLIENT_SECRET;
const GITHUB_WEBHOOK_SECRET = process.env.GITHUB_WEBHOOK_SECRET;
const GITHUB_PRIVATE_KEY = fs.readFileSync(process.env.GITHUB_PRIVATE_KEY_PATH, 'utf8');

mongoose.connect(MONGO_URI)
  .then(() => console.log('MongoDB connected successfully!'))
  .catch(err => {
    console.error('MongoDB connection error:', err);
    process.exit(1); // Exit the process if MongoDB connection fails
  });

app.use(session({
  name: 'gas-cost-monitor.sid',
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: false,
  cookie: { secure: 'auto', httpOnly: true, maxAge: 1000 * 60 * 60 * 24 }
}));

app.get('/auth/github', (req, res) => {
  if (!req.session.user) {
    return res.status(401).send('You must be logged in.');
  }
  const url = `https://github.com/apps/gas-cost-monitor/installations/new`;
  res.redirect(url);
});

app.get('/auth/github/callback', (req, res) => {
  const { installation_id } = req.query;

  if (!req.session.user) {
    return res.status(401).send('Authentication Error');
  }
  console.log(`User ${req.session.user.walletAddress} installed app with installation_id: ${installation_id}`);

  // Redirect user back to the frontend dashboard
  res.redirect('http://localhost/dashboard');
});

app.post('/webhooks/github', (req, res) => {
  const event = req.headers['x-github-event'];
  const payload = req.body;

  if (event === 'pull_request') {
    if (payload.action === 'opened'|| payload.action === 'synchronize') {
      console.log('Pull request opened or updated!');
      handlePullRequest(payload);
    }
  }

  res.status(200).send('Event received');
});

async function handlePullRequest(payload) {
  const installationId = payload.installation.id;
  const repoOwner = payload.repository.owner.login;
  const repoName = payload.repository.name;
  const prNumber = payload.pull_request.number;

  try {
    // TODO: Authenticate as the GitHub App
    // TODO: Get the code from the PR
    // TODO: Run the gas analysis (for now, we'll just mock it)
    // TODO: Post a comment back to the PR

    const commentBody = `
### ⛽ Gas Cost Report ⛽

This is a mock report. In the future, this will contain a detailed analysis of your changes.

*A real analysis is coming soon!*
        `;
        
    // This is a placeholder for the real logic
    console.log(`Would post to ${repoOwner}/${repoName}#${prNumber}: ${commentBody}`);

  } catch (error) {
    console.error('Failed to handle pull request:', error);
  }
}


app.get('/api/auth/nonce', async (req, res) => {
  req.session.nonce = Math.random().toString(36).substring(2);
  await req.session.save();
  res.send(req.session.nonce);
});

app.post('/api/auth/verify', async (req, res) => {
  try {
    const { message, signature } = req.body;
    const siweMessage = new SiweMessage(message);

    const fields = await siweMessage.verify({ signature, nonce: req.session.nonce });

    let user = await User.findOne({ walletAddress: fields.data.address.toLowerCase() });

    if (!user) {
      user = new User({ walletAddress: fields.data.address.toLowerCase(), nonce: fields.data.nonce });
      await user.save();
    }

    req.session.user = { walletAddress: user.walletAddress };
    await req.session.save();

    res.status(200).json({ walletAddress: user.walletAddress });
  } catch (error) {
    console.error('Verification failed:', error);
    res.status(401).json({ message: 'Verification failed' });
  }
});

app.get('/api/auth/me', (req, res) => {
  if (req.session.user) {
    res.json(req.session.user);
  } else {
    res.status(401).json({ message: 'Not authenticated' });
  }
});



app.get("/", (req, res) => {
  res.status(200).json({ message: "Gas Cost Monitor API is running!" });
});

const authenticateToken = (req, res, next) => {
  const authHeader = req.headers['authorization'];
  const token = authHeader && authHeader.split(' ')[1];

  if (token == null) {
    return res.status(401).send('Authentication token required.');
  }

  jwt.verify(token, JWT_SECRET, (err, user) => {
    if (err) {
      return res.status(403).send('Invalid or expired token.');
    }
    req.user = user;
    next();
  });
};

app.post('/api/reports', async (req, res) => {
  try {
    const gasReport = {
      ...req.body,
      tenantId: req.body.tenantId.toLowerCase(), // normalize case
    };

    const newGasReport = new GasReport(gasReport);
    await newGasReport.save();

    console.log('Gas Report saved to MongoDB:', newGasReport);
    res.status(201).json({
      message: 'Gas Report received and saved successfully!',
      data: newGasReport,
    });
  } catch (error) {
    console.error('Error saving gas report to MongoDB:', error);
    res.status(500).json({ error: error.message });
  }
});

app.get('/api/reports', async (req, res) => {
  if (!req.session.user) {
    return res.status(401).json({ message: 'Not authenticated' });
  }
  try {
    console.log(GasReport);
    const reports = await GasReport.find({ tenantId: req.session.user.walletAddress.toLowerCase() }).sort({ timestamp: -1 });
    res.json(reports);
  } catch (error) {
    console.error('Error fetching gas reports:', error);
    res.status(500).send('Server Error');
  }
});

app.post('/api/auth/web3-login', (req, res) => {
  const { walletAddress, signature, message } = req.body;

  if (!walletAddress || !signature || !message) {
    return res.status(400).json({ error: 'Wallet address, signature, and message are required.' });
  }

  try {

    // Verify the signature. This is the core security step.
    const recoveredAddress = ethers.verifyMessage(message, signature);

    if (recoveredAddress.toLowerCase() === walletAddress.toLowerCase()) {
      // Signature is valid. Generate a JWT.
      const payload = {
        tenantId: walletAddress,
        issuedAt: Date.now(),
      };
      const token = jwt.sign(payload, JWT_SECRET, { expiresIn: '1h' });
      
      // Respond with the JWT and the tenantId
      res.status(200).json({ token, tenantId: walletAddress });
    } else {
      // Signature is invalid
      res.status(401).json({ error: 'Invalid signature.' });
    }

  } catch (error) {
    console.error('Signature verification failed:', error);
    res.status(500).json({ error: 'An error occurred during authentication.' });
  }
});

app.get('/api/protected-data', authenticateToken, (req, res) => {
  res.status(200).json({
    message: `Hello, ${req.user.tenantId}! This is protected data.`,
    user: req.user,
  });
});

app.listen(PORT, '0.0.0.0', () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
