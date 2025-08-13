import { useState, useEffect } from 'react';
import axios from 'axios';
import { ethers } from 'ethers';
import { SiweMessage } from 'siwe';
import './App.css';

// The main application component
function App() {
  const [user, setUser] = useState(null);

  // Check if the user is already logged in on component mount
  useEffect(() => {
    const checkUser = async () => {
      try {
        const { data } = await axios.get('/api/auth/me');
        setUser(data);
      } catch (error) {
        setUser(null);
      }
    };
    checkUser();
  }, []);

  const handleLogin = async () => {
    if (!window.ethereum) {
      alert('Please install MetaMask!');
      return;
    }

    try {
      // Request account access
      await window.ethereum.request({ method: 'eth_requestAccounts' });

      // Initialize ethers v5 provider
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();

      // Wait for signer to have address
      const address = await signer.getAddress();
      const { chainId } = await provider.getNetwork();

      // Get nonce from backend
      const { data: nonce } = await axios.get('/api/auth/nonce');

      console.log('provider: ', provider);
      console.log('signer: ', signer);
      console.log('address: ', address);
      console.log('chainId: ', chainId);
      console.log('nonce: ', nonce);

      // Create SIWE message
      const message = new SiweMessage({
        domain: window.location.host,
        address,
        statement: 'Sign in with Ethereum to the Gas Cost Monitor.',
        uri: window.location.origin,
        version: '1',
        chainId,
        nonce,
      });

      console.log("--------------");
      console.log(message);

      // Sign the message
      const signature = await signer.signMessage(message.prepareMessage());

      // Verify on backend
      const { data: verifiedUser } = await axios.post('/api/auth/verify', {
        message: message.prepareMessage(),
        signature,
      });

      setUser(verifiedUser);
    } catch (error) {
      console.error('Login failed:', error);
      setUser(null);
    }
  };


  // Render the Login view or the Dashboard view based on user state
  if (!user) {
    return <LoginScreen onLogin={handleLogin} />;
  }

  return <DashboardScreen user={user} />;
}

// The login screen component
function LoginScreen({ onLogin }) {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Welcome to Gas Cost Monitor</h1>
        <p>Connect your wallet to continue</p>
        <button onClick={onLogin} className="login-button">
          Connect Wallet & Sign In
        </button>
      </header>
    </div>
  );
}

// The dashboard component
function DashboardScreen({ user }) {
  const [reports, setReports] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get('/api/reports')
      .then(response => {
        setReports(response.data);
        setLoading(false);
      })
      .catch(err => {
        setError('Failed to load reports. Ensure your TENANT_ID matches your wallet address.');
        setLoading(false);
      });
  }, []);

  console.log(reports);
  
  return (
    <div className="App">
      <header className="App-header">
        <h1>Gas Cost Monitor Dashboard</h1>
        <p>Welcome, <span className="wallet-address">{user.walletAddress}</span></p>
      </header>
      <main>
        {loading && <div>Loading reports...</div>}
        {error && <div style={{ color: 'orange' }}>{error}</div>}
        {!loading && !error && (
          <table>
            <thead>
              <tr>
                <th>Contract</th>
                <th>Function</th>
                <th>Gas Used</th>
                <th>Timestamp</th>
              </tr>
            </thead>
            <tbody>
              {reports.length > 0 ? (
                reports.map(report => (
                  <tr key={report._id}>
                    <td>{report.contractName}</td>
                    <td>{report.method}</td>
                    <td>{report.gasUsed.toLocaleString()}</td>
                    <td>{new Date(report.timestamp).toLocaleString()}</td>
                  </tr>
                ))
              ) : (
                <tr>
                  <td colSpan="4">No reports found for your account.</td>
                </tr>
              )}
            </tbody>
          </table>
        )}
      </main>
    </div>
  );
  
}

export default App;