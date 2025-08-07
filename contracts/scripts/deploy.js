
const hre = require("hardhat");

async function main() {
    const [deployer] = await hre.ethers.getSigners();

    console.log("Deploying contract with the account:", deployer.address);
    const balance = await hre.ethers.provider.getBalance(deployer.address);
    console.log("Account balance:", hre.ethers.formatEther(balance), "ETH");

    const ContractName = await hre.ethers.getContractFactory("MyContract");
    const contractName = await ContractName.deploy(100);

    await contractName.waitForDeployment();

    console.log("Contract deployed to: ", await contractName.getAddress());

    const receipt = await contractName.deploymentTransaction().wait();
    console.log("Transaction hash:", receipt.hash);
    console.log("Gas used:", receipt.gasUsed.toString());
    console.log("Block number:", receipt.blockNumber);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });