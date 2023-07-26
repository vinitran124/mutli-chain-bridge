// @ts-nocheck
import { useEffect, useState } from 'react';
import Web3, { Contract } from 'web3';
import { Chain } from '../screen/bridge.screen';
import BridgeRouterAbi from '../const/bridgerouter.json';

export const useContract = (
  walletAddress: string,
  contractAddress: any,
  contractAbi: any = BridgeRouterAbi.abi,
) => {
  const [contract, setContract] = useState<Contract<any> | undefined>();
  const [currentChainId, setCurrentChainId] = useState<bigint | undefined>();
  const web3 = new Web3(window.ethereum);

  const initializeContract = async () => {
    if (!walletAddress || !contractAddress) {
      return;
    }

    // Check if Web3 is injected by MetaMask or any other Ethereum provider
    if (window.ethereum) {
      // Initialize Web3 with the current provider
      const web3 = new Web3(window.ethereum);
      try {
        const networkId = await web3.eth.net.getId();

        // Use the network ID to create an instance of the ERC20 contract
        const erc20Contract = new web3.eth.Contract(
          contractAbi,
          contractAddress,
        );
        // Set the contract instance in the context
        setContract(erc20Contract);
        console.log('Connected to ERC20 network:', networkId, erc20Contract);
      } catch (error) {
        console.error('Error connecting to Ethereum network:', error);
      }
    } else {
      console.error('Web3 provider not found. Please install MetaMask.');
    }
  };

  useEffect(() => {
    initializeContract();
    getChainId();
  }, [contractAbi, contractAddress, walletAddress]);

  useEffect(() => {
    if (window.ethereum) {
      window.ethereum.on('chainChanged', () => {
        console.log('on change network');
        getChainId();
        initializeContract();
      });
      window.ethereum.on('accountsChanged', () => {
        getChainId();
        initializeContract();
      });
    }
  }, []);

  const getTokenAvailableInPool = async (tokenAddress: any) => {
    if (contract) {
      try {
        const web3 = new Web3(window.ethereum);
        const balance = await contract.methods
          .getAmountTokenInPool(tokenAddress)
          .call();

        const bInt = web3.utils.fromWei(balance, 'ether');
        console.log('token in pool', bInt);
        return bInt;
      } catch (error) {
        console.error('Error checking token balance:', error);
      }
    }
  };

  const getTransferContractAdd = async (tokenAdd: string) => {
    if (contract) {
      try {
        const contractAdd = await contract.methods.bridgePools(tokenAdd).call();

        return contractAdd as string;
      } catch (error) {
        console.error('Error get transfer contract add:', error);
      }
    }
  };

  const getChainId = async () => {
    try {
      const web3 = new Web3(window.ethereum);

      const chainId = await web3.eth.getChainId();
      console.log('current chain id', chainId);
      setCurrentChainId(chainId);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const changeNetwork = async (chain: Chain) => {
    if (window.ethereum) {
      try {
        const web3 = new Web3(window.ethereum);
        await window.ethereum.request({
          method: 'wallet_switchEthereumChain',
          params: [{ chainId: web3.utils.toHex(BigInt(chain.chainId)) }],
        });
      } catch (error) {
        console.error(error);
      }
    }
  };

  const tranfer = async () => {
    if (contract) {
      try {
        const web3 = new Web3(window.ethereum);
        const balance = await contract.methods
          .transfer(
            '0xD1C80e25CDb409b3F3cB9340a8e35f511A7EbE1F',
            web3.utils.toWei('1000', 'ether'),
          )
          .send({ from: '0x06D74c06196a814F56d552aF83F893689e5e7eF8' });
        const bInt = web3.utils.fromWei(balance, 'ether');
      } catch (error) {
        console.error('Error checking token balance:', error);
      }
    }
  };

  const deposit = async (tokenAdd: string, amount: string) => {
    if (contract) {
      try {
        return contract.methods
          .deposit(tokenAdd, web3.utils.toWei(amount, 'ether'))
          .send({ from: walletAddress });
      } catch (error) {
        console.error('Error in deposit:', error, tokenAdd, amount);
      }
    }
  };

  const getAmountSwap = async (
    amount: string,
    tokenIn: string,
    tokenOut: string,
  ) => {
    if (contract) {
      try {
        const amountOut = (await contract.methods
          .getAmountOut(web3.utils.toWei(amount, 'ether'), tokenIn, tokenOut)
          .call()) as bigint;

        return web3.utils.fromWei(amountOut, 'ether').toString();
      } catch (error) {
        console.error(
          'Error in getAmountt:',
          error,
          web3.utils.toWei(amount, 'ether'),
          tokenIn,
          tokenOut,
        );
      }
    }
  };

  const swapForToken = async (
    amountIn: string,
    amountOut: string,
    tokenIn: string,
    tokenOut: string,
  ) => {
    if (contract) {
      try {
        return contract.methods
          .swapExactETHForTokens(
            web3.utils.toWei(+amountOut * 0.9, 'ether'),
            [tokenIn, tokenOut],
            walletAddress,
            Date.now() + 300,
          )
          .send({
            from: walletAddress,
            value: web3.utils.toWei(amountIn, 'ether'),
          });
      } catch (error) {
        console.error('Error in getAmountt:', error);
      }
    }
  };

  const swapForCoin = async (
    amountIn: string,
    amountOut: string,
    tokenIn: string,
    tokenOut: string,
  ) => {
    if (contract) {
      try {
        return contract.methods
          .swapExactTokensForETH(
            web3.utils.toWei(+amountIn, 'ether'),
            web3.utils.toWei(+amountOut * 0.9, 'ether'),
            [tokenIn, tokenOut],
            walletAddress,
            Date.now() + 300,
          )
          .send({
            from: walletAddress,
          });
      } catch (error) {
        console.error('Error in getAmountt:', error);
      }
    }
  };
  const addLiquidity = async (tokenAdd: string, amount: string) => {
    if (contract) {
      try {
        await contract.methods
          .addLiquidity(tokenAdd, web3.utils.toWei(amount, 'ether'))
          .send({ from: walletAddress });
      } catch (error) {
        console.error('Error in deposit:', error, tokenAdd, amount);
      }
    }
  };

  return {
    getTokenAvailableInPool,
    tranfer,
    currentChainId,
    changeNetwork,
    getTransferContractAdd,
    deposit,
    getAmountSwap,
    swapForToken,
    swapForCoin,
    addLiquidity,
    contract,
  };
};
