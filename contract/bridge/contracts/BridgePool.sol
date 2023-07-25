// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

import "./interfaces/IBridgePool.sol";


contract BridgePool  is IBridgePool, Ownable {
    using Counters for Counters.Counter;
    using SafeERC20 for IERC20;

    Counters.Counter private _addLiquidityIds;

    IERC20 public TOKEN;
    mapping (address => uint256[]) public liquidityIds;
    mapping (uint256 => LiquidityTransaction) public liquidityTx;

    uint256[4] private rewardRate = [10, 15, 20, 25]; // rate = rewardRate /1000

    struct LiquidityTransaction {
        address user;
        uint256 amount;
        uint256 startTime;
    }

    constructor() {}

    function addLiquidity(uint256 _amount) external onlyOwner{
        TOKEN.safeTransferFrom(tx.origin, address(this), _amount);
        uint256 addLiquidityId = _addLiquidityIds.current();
        _addLiquidityIds.increment();
        liquidityIds[tx.origin].push(addLiquidityId);
        liquidityTx[addLiquidityId] = LiquidityTransaction(tx.origin,_amount, block.timestamp);
    }

    function removeLiquidity(uint256 _amount, uint256 _liquidityId) external onlyOwner returns(uint256) {
        require(liquidityTx[_liquidityId].user == tx.origin, "liquidityId does not match with user");
        require(_amount <= liquidityTx[_liquidityId].amount, "amount must be less than user's liquidity amount");    
        liquidityTx[_liquidityId].amount -= _amount;
        // Assume 1 second = 1 block  .31536000 is the total block that rendered in 1 year.
        uint256 reward = _amount * (block.timestamp - liquidityTx[_liquidityId].startTime) * rewardRate[_getRateReward(liquidityTx[_liquidityId].startTime)] / (1000 * 31536000);
        TOKEN.transfer(msg.sender, _amount + reward);
        return _amount + reward;
    }

    function deposit(uint256 _amount) external onlyOwner{
        TOKEN.safeTransferFrom(tx.origin, address(this), _amount);
    }

    function withdraw(uint256 _amount) external onlyOwner {
        TOKEN.transfer(msg.sender, _amount);
    }

    function _getRateReward(uint256 _startTime) private view returns(uint256) {
        uint256 time = block.timestamp - _startTime;
        // 1s = 1 block
        // 3 months = 7884000
        if (time < 7884000) {
            return 0;
        }
        // 6 months = 15768000
        if (time < 15768000) {
            return 1;
        }
        // 9 months = 23652000
        if (time < 23652000) {
            return 2;
        }
        return 3;
    }

    function initialize(address _token) external onlyOwner{
        TOKEN = IERC20(_token);
    }

    function setRate(uint256[4] calldata _rewardRate) external onlyOwner {
        rewardRate = _rewardRate;
    }
}