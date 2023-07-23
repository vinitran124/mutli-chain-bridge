import { useState } from "react"
import { FaEthereum } from "react-icons/fa6"
import { RxTriangleDown } from "react-icons/rx"
import { Chain } from "../../screen/bridge.screen";

interface Props {
    value: Chain;
    options: Chain[];
    onPress: (chain: Chain) => void;
}

export const BridgeDropdown = ({ value, options, onPress }: Props) => {
    const [isOpen, setOpen] = useState(false)

    const onPressItem = (chain: Chain) => {
        setOpen(false);
        onPress(chain)
    }

    return (
        <div className=" relative">
            <div className=" w-[20rem] rounded border border-slate-400 bg-slate-800 px-3 py-4 flex flex-row items-center justify-between cursor-pointer" onClick={() => setOpen(!isOpen)}>
                <div className=" flex flex-row items-center">
                    <p className=" ml-2 font-medium text-lg text-white">
                        {value.name}
                    </p>

                </div>
                <RxTriangleDown color="white" />
            </div>
            {
                isOpen && <div className=" absolute w-[20rem] rounded border border-slate-400 bg-slate-700 max-h-48 overflow-y-scroll cursor-pointer">
                    {
                        options.map(chain => {
                            return (
                                <div className=" flex flex-row items-center px-3 py-4 cursor-pointer" onClick={() => onPressItem(chain)}>
                                    <p className=" ml-2 font-medium text-lg text-white">
                                        {chain.name}
                                    </p>
                                </div>
                            )
                        })
                    }
                </div>
            }
        </div>
    )
}