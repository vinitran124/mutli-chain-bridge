// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IWETH.sol";
import "hardhat/console.sol";

contract WETH is IWETH, ERC20 {
    address public bridgePool;

    constructor(string memory _name, string memory _symbol) ERC20(_name, _symbol) {}
    
    function changeBridgePool(address _bridgePool) external  { 
        bridgePool = _bridgePool;
    }

    function deposit() external payable {
        require(msg.value > 0, "You need to send some ether");
        _mint(tx.origin, msg.value);
    }

    function withdraw(uint256 _amount) external {
        _burn(msg.sender, _amount);
        (bool success, ) = payable(msg.sender).call{value: _amount}("");
        require(success, "ETHW: Failed to send Ether");
    }
}