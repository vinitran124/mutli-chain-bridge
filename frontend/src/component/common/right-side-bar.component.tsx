import React, { useEffect, useState } from 'react';
import './right-side-bar.css';
import { Observable } from 'rxjs';
import { MdOutlineKeyboardDoubleArrowRight } from 'react-icons/md'
import { MetaMaskSDK, MetaMaskSDKOptions } from '@metamask/sdk';
import { useAppDispatch, useAppSelector } from '../../hook/store.hook';
import { setWalletAddress } from '../../store/wallet.slice';
import detectEthereumProvider from '@metamask/detect-provider';
import { formatBalance } from '../../utils';
import { formatAccount } from '../../helper/account.helper';

interface SidebarProps {
    listener: Observable<boolean>
}

const metaMaskOption: MetaMaskSDKOptions = {
    dappMetadata: {},
    checkInstallationImmediately: true
}

export const Sidebar = ({ listener }: SidebarProps) => {
    const [balance, setBalance] = useState(formatBalance('0'))
    const jazzicon = require('@metamask/jazzicon')
    const MMSDK = new MetaMaskSDK(metaMaskOption);
    const ethereum = MMSDK.getProvider(); // You can also access via window.ethereum
    const walletAddress = useAppSelector(state => state.address)
    const dispatch = useAppDispatch();
    const [isOpen, setIsOpen] = useState(false);

    const toggleSidebar = () => {
        setIsOpen(!isOpen);
    };

    useEffect(() => {
        const listen = listener.subscribe((value) => {
            setIsOpen(value)
        })

        return () => {
            listen.unsubscribe();
        }
    }, [])

    useEffect(() => {
        if (walletAddress) {
            window.ethereum!.request({   /* New */
                method: "eth_getBalance",                                      /* New */
                params: [walletAddress, "latest"],                               /* New */
            }).then((balance => setBalance(formatBalance(balance as string || '0'))))
                .catch(error => { })
        }
    }, [walletAddress])

    const openMetamaskWallet = async () => {
        const provider = await detectEthereumProvider({ silent: true })

        if (!Boolean(provider)) {
            const confirmation = window.confirm(
                'MetaMask is not installed. Would you like to install it?'
            );
            if (confirmation) {
                window.open(
                    'https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn',
                    '_blank'
                );
            }
        }

        window.ethereum?.request({   /* New */
            method: "eth_requestAccounts",                 /* New */
        }).then((accounts: any) => {
            dispatch(setWalletAddress(accounts?.[0]))
            /* New */
        }).catch(error => {
            console.log(error)
        })
    }

    const handleDisconnect = () => {
        dispatch(setWalletAddress(''))
    }

    return (
        <div className={`sidebar ${isOpen ? 'open' : ''} cursor-pointer`} onClick={() => setIsOpen(false)}>
            <div className='sidebar-pad-left mt-1'>
                <MdOutlineKeyboardDoubleArrowRight color='white' className='w-6 h-6 ml-2 mt-2 cursor-pointer' onClick={toggleSidebar} />
            </div>
            <div className='sidebar-inner'>
                {walletAddress ?
                    <>
                        <div className=' flex flex-row items-center justify-between px-3 py-2'>
                            <div className=' flex flex-row'>
                                {/* {jazzicon(20, parseInt(walletAddress.slice(2, 10), 16))} */}
                                <div className="flex justify-center items-center bg-orange-600 text-white py-2 rounded-full px-3 mr-4 cursor-pointer max-w-[200px] truncate font-medium"><p className=" truncate">{formatAccount(walletAddress)}</p></div>
                            </div>
                            {/* <div>
                                <FiLogOut color='white' size={'1.5rem'} onClick={handleDisconnect} className='cursor-pointer' />
                            </div> */}

                        </div>
                        <div className=' text-white text-4xl font-medium px-3 mt-3'>${balance}</div>
                    </> : <>
                        <div className='text-lg mx-4 text-white font-medium mt-2 mb-7'>Connect a wallet</div>
                        <div className='flex flex-row items-center px-5 py-4 bg-slate-900 rounded-2xl mx-3' onClick={openMetamaskWallet}>
                            <img src='https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/MetaMask_Fox.svg/800px-MetaMask_Fox.svg.png' className=' rounded-md w-12 h-12 bg-white' />
                            <div className=' text-xl font-bold text-white ml-4'>MetaMask</div>
                        </div>
                    </>
                }

            </div>
        </div>
    );
};

