const { expect } = require("chai");
const { ethers } = require("hardhat");
const { mine } = require("@nomicfoundation/hardhat-network-helpers");

const fromWei = (value) => {
  return ethers.formatEther(value);
}

const toWei = (value) => {
  return ethers.parseUnits(value.toString(), "ether");
}

describe("Bridge", function () {
  let deployer, addr1, addr2, vini, bridgeRouter, bridgePool, bridgePoolWETH
  const nativeToken = "0x0000000000000000000000000000000000000000"

  beforeEach(async function () {
    const BRIDGEPOOL = await ethers.getContractFactory("BridgePool");
    const TOKEN = await ethers.getContractFactory("Token");
    const WETH = await ethers.getContractFactory("WETH");
    const BRIDGEROUTER = await ethers.getContractFactory("BridgeRouter");

    [deployer, addr1, addr2] = await ethers.getSigners();
    
    // vini = await TOKEN.deploy("Vini Tran", "VINI", toWei(1000000000));
    // weth = await WETH.deploy("Wrapper Ether", "WETH");
    bridgeRouter = await BRIDGEROUTER.deploy("0xB378B8c7311898dbE9220aF2c1555369c892eeC3");

    await bridgeRouter.connect(deployer).createBridgePool("0x5cd4a973196cd838CD61d551fbb34E0125bEb82B");
    // const bridgePoolAddr = await bridgeRouter.bridgePools(vini.target);
    // bridgePool = await BRIDGEPOOL.attach(bridgePoolAddr);

    // await bridgeRouter.connect(deployer).createBridgePool(weth.target);
    // const bridgePoolWETHAddr = await bridgeRouter.bridgePools(weth.target);
    // bridgePoolWETH = await BRIDGEPOOL.attach(bridgePoolWETHAddr);

    // await weth.connect(deployer).changeBridgePool(bridgePoolWETH.target);
    // expect(await weth.bridgePool()).to.equal(bridgePoolWETH.target);
    
    // vini.connect(deployer).transfer(addr1.address, toWei(1000000))
    // await vini.connect(addr1).approve(bridgePool.target, toWei(1000000000));

    // vini.connect(deployer).transfer(addr2.address, toWei(1000000))
    // await vini.connect(addr2).approve(bridgePool.target, toWei(1000000000));

    // await weth.connect(addr1).approve(bridgePoolWETH.target, toWei(10000000000));

    // await bridgeRouter.connect(deployer).grantRole(deployer.address);
    // // await bridgeRouter.connect(deployer).revokeRole(deployer.address);
  });

  describe("Deployment", function () {
    it("Add Liquidity", async function () {
      const amountAddLiquidity = 10000
      await bridgeRouter.connect(addr1).addLiquidity(vini.target, toWei(amountAddLiquidity));
      expect(await vini.balanceOf(bridgePool.target)).to.equal(toWei(amountAddLiquidity));
    
      // const balance0ETH = await ethers.provider.getBalance(addr1.address);
      // console.log(fromWei(balance0ETH));
    });

    it("Add Liquidity with native token", async function () {
      const amountETHAddLiquidity = 100;
      await bridgeRouter.connect(addr1).addLiquidity(nativeToken, toWei(amountETHAddLiquidity), { value: toWei(amountETHAddLiquidity) });
    });

    it("Remove Liquidity", async function () {
      const amountAddLiquidity = 10000
      const amountRemoveLiquidity = 1000
      const blockTime = 7884000 -1
      const totalBlockInYear = 31536000
      const totalRemove = amountRemoveLiquidity + amountRemoveLiquidity * ((blockTime +1) / totalBlockInYear) * (15/1000)
      
      await bridgeRouter.connect(addr1).addLiquidity(vini.target, toWei(amountAddLiquidity));
      const beforeAmountVini = await vini.balanceOf(bridgePool.target)
      expect(beforeAmountVini).to.equal(toWei(amountAddLiquidity));

      await mine(blockTime);
      await bridgeRouter.connect(addr1).removeLiquidity(vini.target, toWei(amountRemoveLiquidity), 0);
      const afterAmountVini = await vini.balanceOf(bridgePool.target)
      expect(fromWei(beforeAmountVini - afterAmountVini)).to.equal(String(totalRemove));
      expect(fromWei(afterAmountVini)).to.equal(fromWei(await bridgeRouter.getAmountTokenInPool(vini.target)));
    });

    it("Remove Liquidity with native token", async function () {
      const blockTime = 7884000 -1
      const amountETHAddLiquidity = 100;
      const totalBlockInYear = 31536000

      await bridgeRouter.connect(addr1).addLiquidity(nativeToken, toWei(amountETHAddLiquidity), { value: toWei(amountETHAddLiquidity) });
      const beforeAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const beforeAmountETHAddr1 = await ethers.provider.getBalance(addr1.address);

      await mine(blockTime);

      const amountRemoveLiquidity = 10;

      const totalRemove = amountRemoveLiquidity + amountRemoveLiquidity * ((blockTime +1) / totalBlockInYear) * (15/1000)

      await bridgeRouter.connect(addr1).removeLiquidity(nativeToken, toWei(amountRemoveLiquidity), 0);
      const afterAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const afterAmountETHAddr1 = await ethers.provider.getBalance(addr1.address); 

      expect(Number(fromWei(beforeAmountETHInPool - afterAmountETHInPool))).to.equal(totalRemove);
      expect(Number(fromWei(afterAmountETHAddr1 - beforeAmountETHAddr1)).toFixed(2)).to.equal(totalRemove.toFixed(2));
    });

    it("deposit", async function () {
      const amountAddLiquidity = 10000
      const amountDeposit = 10
      

      await bridgeRouter.connect(addr1).addLiquidity(vini.target, toWei(amountAddLiquidity));
      const beforeAmountViniInPool = await vini.balanceOf(bridgePool.target)
      const beforeAmountViniAddr2 = await vini.balanceOf(addr2.address)

      await bridgeRouter.connect(addr2).deposit(vini.target, toWei(amountDeposit));
      const afterAmountViniInPool = await vini.balanceOf(bridgePool.target)
      const afterAmountViniAddr2 = await vini.balanceOf(addr2.address)

      expect(Number(fromWei(afterAmountViniInPool - beforeAmountViniInPool))).to.equal(amountDeposit);
      expect(Number(fromWei(beforeAmountViniAddr2 - afterAmountViniAddr2))).to.equal(amountDeposit);
    })

    it("deposit native token", async function () {
      const amountETHAddLiquidity = 100
      const amountDeposit = 10

      await bridgeRouter.connect(addr1).addLiquidity(nativeToken, toWei(amountETHAddLiquidity), { value: toWei(amountETHAddLiquidity) });
      const beforeAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const beforeAmountETHAddr1 = await ethers.provider.getBalance(addr1.address);

      await bridgeRouter.connect(addr1).deposit(nativeToken, toWei(amountDeposit), { value: toWei(amountDeposit) });
      const afterAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const afterAmountETHAddr1 = await ethers.provider.getBalance(addr1.address);
    
      expect(Number(fromWei(afterAmountETHInPool - beforeAmountETHInPool)).toFixed(2)).to.equal(String(amountDeposit.toFixed(2)));
      expect(Number(fromWei(beforeAmountETHAddr1 - afterAmountETHAddr1)).toFixed(2)).to.equal(String(amountDeposit.toFixed(2)));
    })

    it("withdraw", async function () {
      const amountAddLiquidity = 10000
      const amountWithdraw = 100

      await bridgeRouter.connect(addr1).addLiquidity(vini.target, toWei(amountAddLiquidity));
      const beforeAmountViniInPool = await vini.balanceOf(bridgePool.target)
      const beforeAmountViniAddr2 = await vini.balanceOf(addr2.address)

      await bridgeRouter.connect(deployer).withdraw(vini.target, addr2.address, toWei(amountWithdraw));
      const afterAmountViniInPool = await vini.balanceOf(bridgePool.target)
      const afterAmountViniAddr2 = await vini.balanceOf(addr2.address)

      expect(Number(fromWei(beforeAmountViniInPool - afterAmountViniInPool))).to.equal(amountWithdraw);
      expect(Number(fromWei(afterAmountViniAddr2 - beforeAmountViniAddr2))).to.equal(amountWithdraw);
    });

    it("withdraw with native token", async function () {
      const amountAddLiquidity = 100
      const amountWithdraw = 10
    
      await bridgeRouter.connect(addr1).addLiquidity(nativeToken, toWei(amountAddLiquidity), { value: toWei(amountAddLiquidity)});
      const beforeAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const beforeAmountETHAddr1 = await ethers.provider.getBalance(addr1.address);

      await bridgeRouter.connect(deployer).withdraw(nativeToken, addr1.address, toWei(amountWithdraw));
      const afterAmountETHInPool = await weth.balanceOf(bridgePoolWETH.target)
      const afterAmountETHAddr1 = await ethers.provider.getBalance(addr1.address);

      expect(Number(fromWei(beforeAmountETHInPool - afterAmountETHInPool))).to.equal(amountWithdraw);
      expect(Number(fromWei(afterAmountETHAddr1 - beforeAmountETHAddr1))).to.equal(amountWithdraw);
    })
  });
});
