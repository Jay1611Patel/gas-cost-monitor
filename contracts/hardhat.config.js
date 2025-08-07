require("@nomicfoundation/hardhat-toolbox");
require("hardhat-gas-reporter");
require("dotenv").config({ path: "../.env" })

/** @type import('hardhat/config').HardhatUserConfig */
/** @type import('hardhat/config').HardhatUserConfig */

const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL || "https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID";
const PRIVATE_KEY = process.env.PRIVATE_KEY || "YOUR_PRIVATE_KEY";

module.exports = {
  solidity: "0.8.20", // Ensure this matches your desired Solidity version
  networks: {
    hardhat: {
      // ...
    },
    sepolia: {
      url: SEPOLIA_RPC_URL,
      accounts: [PRIVATE_KEY],
      chainId: 11155111
    },
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
