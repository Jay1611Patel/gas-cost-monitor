const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("MyContract Gas Tests", function () {
  let MyContract;
  let myContract;
  let owner, addr1, addr2;

  beforeEach(async function () {
    [owner, addr1, addr2] = await ethers.getSigners();
    MyContract = await ethers.getContractFactory("MyContract");
    myContract = await MyContract.deploy(100);
    await myContract.waitForDeployment();
  });

  describe("Value Operations", function () {
    it("should measure gas for setUintValue", async function () {
      const tx = await myContract.setUintValue(200);
      const receipt = await tx.wait();
      expect(await myContract.myUintValue()).to.equal(200);

      // Optional: Log gas used
      console.log("setUintValue gas used:", receipt.gasUsed.toString());
    });
  });

  describe("Address Array Operations", function () {
    it("should add address to array", async function () {
      const tx = await myContract.addAddressToArray(addr1.address);
      const receipt = await tx.wait();
      expect(await myContract.addressArray(0)).to.equal(addr1.address);
      console.log("addAddressToArray gas used:", receipt.gasUsed.toString());
    });

    it("should remove last address from array", async function () {
      await myContract.addAddressToArray(addr1.address);
      const tx = await myContract.removeLastAddressFromArray();
      await tx.wait();

      expect(await myContract.getAddressArrayLength()).to.equal(0);
    });
  });

  describe("Mapping Operations", function () {
    it("should add multiple values to mapping (small)", async function () {
      const keys = [1n, 2n, 3n];
      const values = [10n, 20n, 30n];
      const tx = await myContract.addMultipleValuesToMapping(keys, values);
      const receipt = await tx.wait();

      expect(await myContract.myMapping(1n)).to.equal(10n);
      console.log("small mapping gas used:", receipt.gasUsed.toString());
    });

    it("should add multiple values to mapping (large)", async function () {
      const keys = Array.from({ length: 10 }, (_, i) => BigInt(i + 100));
      const values = Array.from({ length: 10 }, (_, i) =>
        BigInt((i + 1) * 100)
      );
      const tx = await myContract.addMultipleValuesToMapping(keys, values);
      const receipt = await tx.wait();

      expect(await myContract.myMapping(105n)).to.equal(600n);
      console.log("large mapping gas used:", receipt.gasUsed.toString());
    });
  });

  describe("Complex Operations", function () {
    it("should perform complex calculation", async function () {
      const tx = await myContract.performComplexCalculation(5n, 7n);
      const receipt = await tx.wait();
      console.log("complex calc gas used:", receipt.gasUsed.toString());
    });

    it("should trigger event and return value", async function () {
      const tx = await myContract.triggerEventAndReturn("Hello Gas!");
      const receipt = await tx.wait();

      await expect(tx).to.emit(myContract, "LogMessage").withArgs("Hello Gas!");
    });
  });

  describe("Conditional Execution", function () {
    it("should increment when true", async function () {
      const initialValue = await myContract.myUintValue();
      const tx = await myContract.conditionalExecution(true);
      const receipt = await tx.wait();

      expect(await myContract.myUintValue()).to.equal(initialValue + 1n);
      console.log("conditional true gas used:", receipt.gasUsed.toString());
    });

    it("should not change value when false", async function () {
      const initialValue = await myContract.myUintValue();
      const tx = await myContract.conditionalExecution(false);
      await tx.wait();

      expect(await myContract.myUintValue()).to.equal(initialValue);
    });
  });
});
