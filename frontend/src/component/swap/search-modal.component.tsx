import { IoMdClose } from 'react-icons/io'
import { BsSearch } from 'react-icons/bs'
import { SearchTokenItem } from './search-token-item.component'
import { useState } from 'react'
import { createPortal } from 'react-dom';

interface SearchModalProps {
    visble: boolean;
    setVisble: (visible: boolean) => void;
}

export const SearchModal = ({ visble, setVisble }: SearchModalProps) => {
    const rootElement = document.getElementById('home') as Element;

    const closeModal = () => {
        setVisble(false)
    }

    return createPortal(
        <div className={`${visble === true ? 'opacity-100 ' : "opacity-0 pointer-events-none"} w-full h-full bg-black/10 z-40 fixed top-0 left-0 ease-in-out duration-300 transition-all`} onClick={closeModal}>
            <div className="w-full h-full bg-slate-900/75 absolute z-10 items-center justify-center flex">
                <div className="rounded-2xl py-4 border-solid border-stone-300/0.5 bg-slate-900 border-2">
                    <div className="flex flex-row justify-between items-center px-4">
                        <div className="text-slate-50 font-normal text-xl">Select a token</div>
                        <div className="cursor-pointer">
                            <IoMdClose color='white' />
                        </div>
                    </div>
                    <div className='px-4 my-2 border-solid border-gray-800 border-b-2 pb-4'>
                        <div className='flex flex-row items-center border-solid border-l-blue-200 rounded-2xl p-2 border-2 w-96'>
                            <BsSearch color='white' />
                            <input type="text" className="bg-transparent ml-2 text-xl outline-none caret-slate-100 text-slate-100 flex-1"
                                placeholder='Search token name' />
                        </div>
                    </div>
                    <div className='pb-4 max-h-64 overflow-auto'>
                        {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((value, index) => {
                            return (
                                <SearchTokenItem />
                            )
                        })}
                    </div>
                </div>
            </div>
        </div>, document.body
    )
}