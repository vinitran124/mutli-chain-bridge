import Web3 from "web3";
import { useAppSelector } from "../hook/store.hook";
import { useEffect, useState } from "react";
import { Data } from "../const/data";
import { BsChevronDown, BsChevronUp } from "react-icons/bs";
import { api } from "../api/api";

export interface Coin {
    name: string;
    address: string;

}

export const FaunetScreen = () => {
    const walletAddress = useAppSelector(state => state.address);
    const [chainId, setCurrentChainId] = useState<string | undefined>()
    const [token, setToken] = useState<Coin | undefined>();
    const [isOpen, setOpen] = useState(false);

    const getChainId = async () => {
        if (!walletAddress) {
            return;
        }

        if (typeof window.ethereum !== 'undefined') {
            try {
                const web3 = new Web3(window.ethereum);

                const chainId = await web3.eth.getChainId();
                setCurrentChainId(chainId.toString());
                setToken(Data.coin[chainId.toString()][0])
            } catch (error) {
                console.error('Error:', error);
            }
        } else {
            console.error('Metamask not found.');
        }
    };

    useEffect(() => {
        if (window.ethereum) {
            window.ethereum.on('chainChanged', () => {
                getChainId();
            });
            window.ethereum.on('accountsChanged', () => {
                getChainId();
            });
        }
        getChainId();
    }, [])

    const onSelectCoin = (coin: Coin) => {
        setToken(coin);
        setOpen(false);
    }

    const onFaucet = () => {
        if (!walletAddress || !chainId) {
            return;
        }

        api.post('/api/v1/faucet', {
            chain_id: chainId,
            token: token?.address,
            user_address: walletAddress
        })
    }

    return (
        <div className="w-full h-full justify-center flex bg-slate-900">
            <div className=" flex flex-row mt-6 items-center">
                <input className=" rounded-xl bg-slate-800 px-3 py-3 flex-1 caret-slate-100 outline-none mr-4 w-96 h-12 text-white" placeholder="Input your address..." />
                <div className="relative">
                    <div className=" flex flex-row items-center rounded-lg bg-slate-800 h-12 mr-3 px-3 cursor-pointer w-20" onClick={() => setOpen(!isOpen)}>
                        <div className=" text-white text-lg mr-2">{token?.name}</div>
                        {(isOpen && !!chainId) ? <BsChevronUp color="white" /> : <BsChevronDown color="white" />}

                    </div>
                    <div>
                        {(isOpen && !!chainId) && <div className=" absolute rounded-lg overflow-hidden bg-slate-800">
                            {
                                Data.coin[chainId].map((coin) => {
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