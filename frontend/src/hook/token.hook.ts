// @ts-nocheck
import { useAppSelector } from './store.hook';
import tAbi from '../const/token.json';
import Web3, { Contract } from 'web3';
import { useEffect, useState } from 'react';

export const useToken = (contractAddress: any, tokenAbi: any = tAbi.abi) => {
  const [contract, setContract] = useState<Contract<any> | undefined>();
  const [tokenBalance, setTokenBalance] = useState<string>();
  const walletAddress = useAppSelector(state => state.address);

  const initializeContract = async () => {
    console.log('init token', contractAddress, walletAddress);
    if (!walletAddress || !contractAddress) {
      return;
    }
    // Check if Web3 is injected by MetaMask or any other Ethereum provider
    if (window.ethereum) {
      // Initialize Web3 with the current provider
      const web3 = new Web3(window.ethereum);
      try {
        // Request access to user accounts if needed
        await window.ethereum?.request({
          /* New */ method: 'eth_requestAccounts' /* New */,
        });
        // Access the connected network ID
        const networkId = await web3.eth.net.getId();

        // Use the network ID to create an instance of the ERC20 contract
        const erc20Contract = new web3.eth.Contract(tokenAbi, contractAddress);
        console.log('init contract', erc20Contract);
        // Set the contract instance in the context
        setContract(erc20Contract);
        console.log('Connected to ERC20 network:', networkId);
      } catch (error) {
        console.error('Error connecting to Ethereum network:', error);
      }
    } else {
      console.error('Web3 provider not found. Please install MetaMask.');
    }
  };

  useEffect(() => {
    initializeContract();
  }, [tokenAbi, contractAddress, walletAddress]);

  useEffect(() => {
    getWalletTokenAmount();
  }, [contract]);

  const getWalletTokenAmount = async () => {
    console.error('get token', contract, contractAddress);
    if (contract) {
      try {
        const web3 = new Web3(window.ethereum);
        const balance = await contract.methods.balanceOf(walletAddress).call();
        const bInt = web3.utils.fromWei(balance, 'ether');
        setTokenBalance(bInt);
        return bInt;
      } catch (error) {
        console.error('Error checking token balance:', error);
      }
    }
  };

  const getAmountCanTranfer = async (contractAdd: string) => {
    if (contract) {
      try {
        const web3 = new Web3(window.ethereum);
        const balance = await contract.methods
          .allowance(walletAddress, contractAdd)
          .call();
        const bInt = web3.utils.fromWei(balance.toString(), 'ether');
        if (bInt[bInt.length - 1] == '.') {
          bInt.substring(0, bInt.length - 2);
        }

        return bInt;
      } catch (error) {
        console.error('Error checking token balance:', error);
      }
    }
  };

  const approveAmountTransfer = async (contractAdd: string) => {
    if (contract) {
      try {
        const web3 = new Web3(window.ethereum);

        await contract.methods
          .approve(contractAdd, web3.utils.toWei(100000000, 'ether'))
          .send({ from: walletAddress });
      } catch (error) {
        console.error('Error checking token balance:', error);
      }
    }
  };

  return {
    getWalletTokenAmount,
    getAmountCanTranfer,
    approveAmountTransfer,
    tokenBalance,
  };
};
