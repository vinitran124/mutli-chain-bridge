// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IBridgePool {
    function addLiquidity(uint256 _amount) external;
    function removeLiquidity(uint256 _amount, uint256 _index) external returns(uint256);
    function deposit(uint256 _amount) external;
    function withdraw(uint256 _amount) external;
    function initialize(address) external;
    function setRate(uint256[4] calldata _rewardRate) external;
}