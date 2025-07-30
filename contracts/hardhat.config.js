require("@nomicfoundation/hardhat-toolbox");
require("hardhat-gas-reporter");

/** @type import('hardhat/config').HardhatUserConfig */
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.20", // Ensure this matches your desired Solidity version
  networks: {
    hardhat: {
      // ...
    },
    // sepolia: { // You'll add this later when setting up external networks
    //   url: process.env.SEPOLIA_RPC_URL,
    //   accounts: [process.env.PRIVATE_KEY]
    // }
  },
  gasReporter: {
    enabled: true, // This enables the gas reporter
    currency: "USD", // Report costs in USD
    outputFile: "./gas-report.json", // Path relative to contracts/
    coinmarketcap: process.env.COINMARKETCAP_API_KEY, // Needs API key for USD conversion
    outputJSON: true,
    token: "ETH", // Specify the token for price fetching (e.g., ETH, BNB, MATIC)
    gasPriceApi:
      "https://api.etherscan.io/api?module=proxy&action=eth_gasPrice", // Or another provider
    // noColors: true,
    // excludeContracts: [],
    // showMethodSig: true,
    // showTimeSpent: true,
  },
};
