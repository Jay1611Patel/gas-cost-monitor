require('dotenv').config({ path: '../../.env' });

const express = require("express");
const mongoose = require('mongoose');

const app = express();
const PORT = process.env.API_PORT || 3000;
const MONGO_URI = process.env.MONGO_URI || 'mongodb://localhost:27017/gas_costs';

mongoose.connect(MONGO_URI)
  .then(() => console.log('MongoDB connected successfully!'))
  .catch(err => {
    console.error('MongoDB connection error:', err);
    process.exit(1); // Exit the process if MongoDB connection fails
  });

// --- GasReport Schema---
const gasReportSchema = new mongoose.Schema({
  contractName: {
    type: String,
    required: true
  },
  method: {
    type: String,
    required: true
  },
  gasUsed: {
    type: Number,
    required: true
  },
  costUsd: {
    type: Number, // Can be null if CoinMarketCap API fails or is not configured
    default: null
  },
  timestamp: {
    type: Date,
    default: Date.now
  },
  // more fields later, like 'network', 'txHash', 'blockNumber', 'commitHash', etc.
});

const GasReport = mongoose.model('GasReport', gasReportSchema);

app.use(express.json());

app.get("/", (req, res) => {
  res.status(200).json({ message: "Gas Cost Monitor API is running!" });
});

app.post("/api/gas-report", async (req, res) => {
  const gasReport = req.body;
  console.log("Received gas report:");
  console.log(JSON.stringify(gasReport, null, 2)); // Pretty print JSON

  try {
    const newGasReport = new GasReport(gasReport);
    await newGasReport.save();

    console.log('Gas Report saved to MongoDB: ', newGasReport);
    res.status(201).json({
      message: 'Gas Report received and saved Successfully!',
      status: 'ok',
      data: newGasReport,
    });
  } catch (error) {
    console.error('Error saving gas report to MongoDB:', error);
    res.status(500).json({
      message: 'Failed to save gas report',
      error: error.message,
    });
  }
  
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
