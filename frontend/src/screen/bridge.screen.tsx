import { FaChevronDown, FaEthereum } from 'react-icons/fa6';
import { useEffect, useState } from 'react';
import { useAppSelector } from '../hook/store.hook';
import { BehaviorSubject } from 'rxjs';
import { BridgeDropdown } from '../component/bridge/bridge-dropdown.component';
import { BsArrowRight } from 'react-icons/bs';
import { useContract } from '../hook/contract.hook';
import { Data } from '../const/data';
import { api } from '../api/api';
import { useToken } from '../hook/token.hook';
import { BridgeStatusModal } from '../component/bridge/bridge-status-modal.component';
import Web3 from 'web3';

interface Props {
  sidebarSubject: BehaviorSubject<boolean>;
}

export interface Chain {
  chainId: string;
  name: string;
  img: string;
  bridgeContractAddress: string;
}

export const Bridge = ({ sidebarSubject }: Props) => {
  const web3 = new Web3(window.ethereum);
  const [modalVisible, setModalVisible] = useState(false);
  const [chainIn, setChainIn] = useState<Chain>(Data.chain[0]);
  const [chainOut, setChainOut] = useState<Chain>(Data.chain[0]);
  const [isOpenToken, setOpenToken] = useState(false);
  const [coin, setCoin] = useState(Data.coin[chainIn.chainId][0]);
  const [amount, setAmount] = useState<string | undefined>();
  const [tokenAvaible, setTokenAvailable] = useState<string | undefined>();
  const [tokenAmountMax, setTokenAmountMax] = useState<string | undefined>();
  const [contractAddress, setContractAddress] = useState<undefined | string>();

  const walletAddress = useAppSelector(state => state.address);
  const {
    currentChainId,
    changeNetwork,
    getTokenAvailableInPool,
    getTransferContractAdd,
    deposit,
  } = useContract(contractAddress);
  const { getWalletTokenAmount, getAmountCanTranfer, approveAmountTransfer } =
    useToken(coin.address);

  useEffect(() => {
    if (currentChainId?.toString() != chainIn.chainId) {
      setContractAddress(undefined);
    }

    setContractAddress(chainIn.bridgeContractAddress);
  }, [currentChainId, chainIn.chainId]);

  useEffect(() => {
    setCoin(Data.coin[chainIn.chainId][0]);
  }, [chainIn.chainId]);

  const onSelectCoin = (coin: any) => {
    setCoin(coin);
    setOpenToken(false);
  };

  const getTokenAmount = () => {
    if (
      !contractAddress ||
      !walletAddress ||
      currentChainId?.toString() != chainIn.chainId
    ) {
      setTokenAmountMax(undefined);
      setTokenAvailable(undefined);
      return;
    }

    const timeout = setTimeout(() => {
      getTokenAvailableInPool(coin.address).then(amount =>
        setTokenAvailable(amount),
      );
      getWalletTokenAmount().then(amount => setTokenAmountMax(amount));
      clearTimeout(timeout);
    }, 3000);
  };
  useEffect(() => {
    getTokenAmount();
  }, [walletAddress, currentChainId, chainIn.chainId]);

  const onBridge = async () => {
    if (!(amount || 0)) {
      return;
    }

    const contractTranferAdd = await getTransferContractAdd(coin.address);
    const allowanceAmount = await getAmountCanTranfer(
      contractTranferAdd as string,
    );
    if (+(allowanceAmount || 0) < +(amount || 0)) {
      await approveAmountTransfer(contractTranferAdd as string);
    }

    api
      .post('/api/v1/bridge', {
        in_chain: chainIn.chainId,
        out_chain: chainOut.chainId,
        token_address: coin.address,
        amount: web3.utils.toWei(+(amount || 0), 'ether'),
        user_address: walletAddress,
      })
      .then((res: any) => {
        if (res.code == 400) {
          return window.alert(res.message);
        }
        deposit(coin.address, amount as string);
      });
  };

  const getSubmitButton = () => {
    if (!walletAddress) {
      return (
        <div
          className=" mt-2 w-full py-3 text-center bg-orange-500/90 rounded-xl text-white font-medium text-xl cursor-pointer"
          onClick={() => {
            sidebarSubject.next(true);
          }}
        >
          Connect wallet
        </div>
      );
    }
    if (currentChainId?.toString() != chainIn.chainId) {
      return (
        <div
          className=" mt-2 w-full py-3 text-center bg-orange-500/90 rounded-xl text-white font-medium text-xl cursor-pointer"
          onClick={() => {
            changeNetwork(chainIn);
          }}
        >
          Change Metamask Network
        </div>
      );
    }

    return (
      <div
        className=" mt-2 w-full py-3 text-center bg-orange-500/90 rounded-xl text-white font-medium text-xl cursor-pointer"
        onClick={onBridge}
      >
        Bridge
      </div>
    );
  };

  const onChangeAmountInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value.replace('/D/g', '');
    setAmount(value);
  };

  const setChainInput = (chain: Chain) => {
    setCoin(Data.coin[chain.chainId][0]);
    setChainIn(chain);
  };

  return (
    <div className="w-full h-full justify-center flex bg-slate-900">
      <div className="border-solid border-2 border-stone-500 rounded-2xl mt-44 h-fit w-[50rem] m-0 p-4">
        <div className=" flex flex-row justify-between items-center">
          <BridgeDropdown
            value={chainIn}
            onPress={setChainInput}
            options={Data.chain}
          />
          <BsArrowRight color="white" size={'2rem'} />
          <BridgeDropdown
            value={chainOut}
            onPress={setChainOut}
            options={Data.chain}
          />
        </div>
        <div className="w-full border border-slate-600 rounded-xl mt-4">
          <div className=" flex flex-row items-center w-full pr-4">
            <div className=" px-3 py-2 text-white mx-3 my-3 border border-slate-600 rounded">
              MAX
            </div>
            <input
              className=" outline-none bg-transparent text-lg text-white caret-slate-100 flex-1"
              onChange={onChangeAmountInput}
              value={amount}
              placeholder="0"
            />
            <div
              className="rounded-full flex flex-row items-center bg-orange-600 p-2 cursor-pointer relative"
              onClick={() => setOpenToken(!isOpenToken)}
            >
              <div className="text-slate-200 font-medium mx-1 cursor-pointer">
                {coin.name}
              </div>
              <FaChevronDown color="white" />
              {isOpenToken && (
                <div className=" absolute -translate-x-1/2 px-6 top-full mt-2 rounded-2xl border bg-orange-500 left-1/2">
                  {Data.coin[chainIn.chainId].map(coin => {
                    return (
                      <div
                        className=" py-2 text-white font-medium text-lg"
                        onClick={() => onSelectCoin(coin)}
                      >
                        {coin.name}
                      </div>
                    );
                  })}
                </div>
              )}
            </div>
          </div>
          {tokenAmountMax && (
            <div className=" text-slate-400 w-full text-right pr-[14px] mb-2">
              Balance: {tokenAmountMax}
            </div>
          )}
        </div>
        {tokenAvaible && (
          <div className=" w-full text-right text-slate-300 pr-4 mt-2">
            Token available in Pool: {tokenAvaible}
          </div>
        )}
        {getSubmitButton()}
      </div>
      <BridgeStatusModal visble={modalVisible} setVisble={setModalVisible} />
    </div>
  );
};
