import Web3 from "web3";
import { useAppSelector } from "../hook/store.hook";
import { useEffect, useState } from "react";
import { Data } from "../const/data";
import { BsChevronDown, BsChevronUp } from "react-icons/bs";
import { api } from "../api/api";
import { Chain } from "./bridge.screen";

export interface Coin {
    name: string;
    address: string;

}

export const FaunetScreen = () => {
    const [chain, setCurrentChain] = useState(Data.chain[1])
    const [token, setToken] = useState<Coin>(Data.coin['5555'][0]);
    const [isOpenCoin, setOpenCoin] = useState(false);
    const [isOpenChain, setOpenChain] = useState(false);
    const [address, setAddress] = useState<string>();

    const onChangeAddress = (e: React.ChangeEvent<HTMLInputElement>) => {
        setAddress(e.target.value)
    }

    const onSelectCoin = (coin: Coin) => {
        setToken(coin);
        setOpenCoin(false);
    }

    const onSelectChain = (chain: Chain) => {
        setCurrentChain(chain)
        setToken(Data.coin[chain.chainId][0])
    }

    const onFaucet = () => {
        if (!address) {
            return;
        }

        api.post('/api/v1/faucet', {
            chain_id: chain.chainId,
            token: token.address,
            user_address: address
        }).then(res => {
            if (res.status == 200) {
                alert('Success');
            } else alert('Error')
        })
    }

    return (
        <div className="w-full h-full justify-center flex bg-slate-900">
            <div className=" flex flex-row mt-6 items-center">
                <input className=" rounded-xl bg-slate-800 px-3 py-3 flex-1 caret-slate-100 outline-none mr-4 w-96 h-12 text-white" placeholder="Input your address..." onChange={onChangeAddress} />
                <div className="relative">
                    <div className=" flex flex-row items-center rounded-lg bg-slate-800 h-12 mr-3 px-3 cursor-pointer w-36" onClick={() => setOpenChain(!isOpenChain)}>
                        <div className=" text-white text-lg mr-2">{chain?.name}</div>
                        {(isOpenChain) ? <BsChevronUp color="white" /> : <BsChevronDown color="white" />}

                    </div>
                    <div>
                        {(isOpenChain) && <div className=" absolute rounded-lg overflow-hidden bg-slate-800">
                            {
                                Data.chain.map((chain) => {
                                    return (
                                        <div className=" flex flex-row items-center bg-slate-800 h-12 px-3 cursor-pointer w-20" onClick={() => onSelectChain(chain)}>
                                            <div className=" text-white text-lg mr-6">{chain.name}</div>
                                        </div>
                                    )
                                })
                            }
                        </div>}
                    </div>
                </div>
                <div className="relative">
                    <div className=" flex flex-row items-center rounded-lg bg-slate-800 h-12 mr-3 px-3 cursor-pointer w-20" onClick={() => setOpenCoin(!isOpenCoin)}>
                        <div className=" text-white text-lg mr-2">{token?.name}</div>
                        {(isOpenCoin) ? <BsChevronUp color="white" /> : <BsChevronDown color="white" />}

                    </div>
                    <div>
                        {(isOpenCoin) && <div className=" absolute rounded-lg overflow-hidden bg-slate-800">
                            {
                                Data.coin[chain.chainId].map((coin) => {
                                    return (
                                        <div className=" flex flex-row items-center bg-slate-800 h-12 px-3 cursor-pointer w-20" onClick={() => onSelectCoin(coin)}>
                                            <div className=" text-white text-lg mr-6">{coin.name}</div>
                                        </div>
                                    )
                                })
                            }
                        </div>}
                    </div>
                </div>
                <div className=" bg-orange-600 h-12 justify-center items-center rounded-xl py-2 px-2 text-center text-white font-medium text-lg cursor-pointer" onClick={onFaucet}>
                    Give me token
                </div>
            </div>
        </div >
    )
}