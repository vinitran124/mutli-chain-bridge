import { useState } from "react";
import { FaEthereum, FaChevronDown } from "react-icons/fa6";
import { SearchModal } from "./search-modal.component";
import { createPortal } from "react-dom";

export const InputToken = () => {
  const [modalVisble, setModalVisble] = useState(false);

  const setModalVisbleState = (visble: boolean) => {
    setModalVisble(visble)
  }

  return (
    <>
      <div className="w-96 bg-slate-800 rounded-xl border-solid border-2 border-slate-700 mt-3">
        <div className="p-4 flex flex-row items-center">
          <input
            type="text"
            placeholder="0"
            className="bg-transparent text-xl outline-none caret-slate-100 text-slate-100"
          />
          <div className="rounded-full flex flex-row items-center bg-blue-700 p-2 cursor-pointer" onClick={() => setModalVisbleState(true)}>
            <FaEthereum color="white" />
            <div className="text-slate-200 font-medium mx-1 cursor-pointer">{"ETH"}</div>
            <FaChevronDown color="white" />
          </div>
        </div>
      </div>
      <SearchModal visble={modalVisble} setVisble={setModalVisbleState} />
    </>
  );
};
