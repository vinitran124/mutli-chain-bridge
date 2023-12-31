// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IWETH {
    function deposit() external payable;
    function withdraw(uint256 _amount) external;
}