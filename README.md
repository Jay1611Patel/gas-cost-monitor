
Contract address
0x0B936ACD811AE89185a390C65350aDE835912710


const MyContract = await ethers.getContractAt("MyContract", "0x0B936ACD811AE89185a390C65350aDE835912710");
const tx = await MyContract.setUintValue(42);
await tx.wait();
console.log("Transaction sent and confirmed!");