import { FaGear } from "react-icons/fa6";
import { CgArrowsExchangeAltV } from 'react-icons/cg';

import { InputToken } from "../component/swap/input-token.component";
import { SearchModal } from "../component/swap/search-modal.component";
import { useState } from "react";
import { useAppSelector } from "../hook/store.hook";
import { BehaviorSubject } from "rxjs";

enum SwapFunction {
  SWAP,
  BUY
}

interface Props {
  sidebarSubject: BehaviorSubject<boolean>
}

export const SwapScreen = ({ sidebarSubject }: Props) => {
  const [funType, setFuncType] = useState(SwapFunction.SWAP);
  const walletAddress = useAppSelector(state => state.address);

  const onPressSwap = () => {
    setFuncType(SwapFunction.SWAP)
  };

  const onPressBuy = () => {
    if (!walletAddress) {
      return sidebarSubject.next(true)
    }
    setFuncType(SwapFunction.BUY)
  }

  const getSubmitButton = () => {
    if (!walletAddress) {
      return (
        <div className=" mt-2 w-full py-3 text-center bg-orange-500/90 rounded-xl text-white font-medium text-xl cursor-pointer" onClick={() => {sidebarSubject.next(true)}}>Connect wallet</div>
      )
    }
  }

  return (
    <div className="w-full h-full items-center justify-center flex bg-slate-900">
      <div className="border-solid border-2 border-stone-500 rounded-2xl self-center w-fit m-0 p-4">
        <div className=" flex flex-row justify-between items-center">
          <div className="flex flex-row">
            <div className={`mr-4 text-lg cursor-pointer ${funType === SwapFunction.SWAP ? 'font-semibold text-slate-200' : 'font-medium text-slate-500'}`} onClick={onPressSwap}>Swap</div>
            <div className={` font-medium text-lg cursor-pointer ${funType === SwapFunction.BUY ? 'font-semibold text-slate-200' : 'font-medium text-slate-500'}`} onClick={onPressBuy}>Buy</div>
          </div>
          <FaGear color="#EDEDED" />
        </div>
        <InputToken />
        <div className="w-full items-center flex justify-center">
          <div className=" p-1 items-center justify-center mt-3 bg-gray-800 w-fit rounded-lg">
            <CgArrowsExchangeAltV color="white" className="w-8 h-8" />
          </div>
        </div>
        <InputToken />
        {getSubmitButton()}
      </div>
    </div>
  );
};
