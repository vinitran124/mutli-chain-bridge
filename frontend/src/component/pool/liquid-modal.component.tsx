import { useEffect, useRef, useState } from "react";
import { createPortal } from "react-dom";
import { BsChevronDown, BsChevronUp } from "react-icons/bs";
import { FaEthereum } from "react-icons/fa6";
import { useAppSelector } from "../../hook/store.hook";
import Web3 from "web3";
import { Data } from "../../const/data";
import { Coin } from "../../screen/faunet.screen";
import { useContract } from "../../hook/contract.hook";

interface Props {
    onCloseModal: () => void;
}
export const LiquidModal = ({ onCloseModal }: Props) => {
    const [token, setToken] = useState<Coin | undefined>();
    const [chainId, setCurrentChainId] = useState<string | undefined>()
    const [isOpen, setOpen] = useState(false);
    const [amount, setAmount] = useState<string | undefined>();

    const modalRef = useRef<any>();
    const walletAddress = useAppSelector(state => state.address);
    const [contractAddress, setContractAddress] = useState<undefined | string>();

    const {addLiquidity} = useContract(contractAddress);

    const onChangeAmount = (e: React.ChangeEvent<HTMLInputElement>) => {
        setAmount(e.target.value)
    }

    const handleCloseModal = (e: any) => {
        if (modalRef.current === e.terget) {
            onCloseModal()
        }
    }

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
                setContractAddress(Data.chain.find(chain => chain.chainId == chainId.toString())?.bridgeContractAddress)
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

    const onSubmit = () => {
        if (!walletAddress || !contractAddress || !token || !amount) {
            return;
        }

        addLiquidity(token.address, amount);
    }

    return createPortal(
        <div ref={modalRef} className={` w-full h-full bg-black/10 z-40 fixed top-0 left-0 ease-in-out duration-300 transition-all`} onClick={handleCloseModal}>
            <div className="w-full h-full bg-slate-900/75 absolute z-10 items-center justify-center flex">
                <div className="rounded-2xl p-4 border-solid border-stone-300/0.5 bg-slate-900 border-2">
                    <div className=" flex flex-row justify-between items-center">
                        <div className=" text-slate-300 font-normal text-lg">Token</div>
                        <div className=" text-slate-300 font-normal text-lg">Deposit amount</div>
                    </div>
                    <div className=" flex flex-row items-center">
                        <div className=" relative">
                            <div className=" mr-4 rounded-xl w-36 bg-slate-800 px-3 py-3 flex flex-row justify-between items-center cursor-pointer" onClick={() => setOpen(!isOpen)}>
                                {/* <FaEthereum color="white" size={"1.5rem"} /> */}
                                <div className=" text-white ml-3 mr-2">{token ? token.name : 'Wrong Metamask network'}</div>
                                {(isOpen && !!chainId) ? <BsChevronUp color="white" size={"1rem"} /> : <BsChevronDown color="white" size={"1rem"} />}

                            </div>
                            {(isOpen && !!chainId) && <div className=" absolute rounded-xl mt-2 overflow-hidden bg-slate-800">
                                {
                                    Data.coin[chainId].map((coin) => {
                                        return (
                                            <div className=" flex flex-row items-center bg-slate-800 h-12 px-3 cursor-pointer w-36" onClick={() => onSelectCoin(coin)}>
                                                <div className=" text-white text-lg mr-6">{coin.name}</div>
                                            </div>
                                        )
                                    })
                                }
                            </div>}
                        </div>
                        <input className=" rounded-xl bg-slate-800 px-3 py-3 flex-1 w-64 caret-slate-100 outline-none " placeholder="0" onChange={onChangeAmount}/>
                    </div>
                    <div className=" w-full bg-orange-600 justify-center items-center rounded-xl py-3 mt-3 text-center text-white font-medium text-lg cursor-pointer" onClick={onSubmit}>
                        Add liquidity
                    </div>
                </div>
            </div>
        </div>
        , document.body
    )
}