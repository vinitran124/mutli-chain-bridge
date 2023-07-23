import React, { useCallback, useEffect, useState } from "react";
import { BsArrowDownUp, BsCheck } from "react-icons/bs";
import { FaArrowDown, FaArrowUp } from "react-icons/fa6";

interface DropdownOption {
    value: string;
    label: string | React.ReactNode
}

interface DropdownProps {
    value: string;
    options: DropdownOption[];
    onSelectValue: (value: string) => void;
}

export const Dropdown = ({ value, onSelectValue, options }: DropdownProps) => {
    const [isOpen, setOpen] = useState(false)
    const [option, setOptioon] = useState(options.find(option => option.value === value))

    const setIsOpen = () => setOpen(!isOpen)

    const renderLabel = (option: DropdownOption) => {
        const onPressOption = () => {
            setOptioon(option);
            onSelectValue(option.value)
            setOpen(false)
        }

        if (!option) {
            return <></>
        }

        if (typeof option.label === 'string') {
            return (<div className={`mr-3 text-white font-medium text-base min-w-[6rem] pl-3 pr-1 py-2 flex flex-row justify-between items-center cursor-pointer`} onClick={onPressOption}><div>{option.label}</div>{option.value === value && <BsCheck />}</div>)
        }

        return <div>{option.label}</div>
    }

    const renderValue = (option?: DropdownOption) => {
        return (<div className={`mr-3 text-white font-medium text-base min-w-[2rem] flex flex-row justify-between`}><div>{option?.label}</div></div>)
    }

    return (
        <>
            <div onClick={setIsOpen} className={`rounded-xl w-fit cursor-pointer flex flex-row items-center justify-between px-3 py-2 ${isOpen ? 'bg-cyan-950' : 'bg-slate-600'}`}>
                {renderValue(option)}
                {isOpen ? <FaArrowUp size={'1rem'} color="white" className=" ml-2" /> : <FaArrowDown size={'1rem'} color="white" className=" ml-2" />}

            </div>
            {isOpen && <div className=" absolute py-2">
                <div className=" bg-slate-800 rounded-2xl border-slate-600 border-2">{
                    options.map(option => {
                        return <>{renderLabel(option)}</>
                    })
                }</div>
            </div>}
        </>
    )
}