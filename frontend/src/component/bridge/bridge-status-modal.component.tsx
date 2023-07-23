import { createPortal } from "react-dom"

interface ModalProps {
    visble: boolean;
    setVisble: (visible: boolean) => void;
}

export const BridgeStatusModal = ({ visble, setVisble }: ModalProps) => {
    return createPortal(
        <div className={`${visble === true ? 'opacity-100 ' : "opacity-0 pointer-events-none"} w-full h-full bg-black/10 z-40 fixed top-0 left-0 ease-in-out duration-300 transition-all`} onClick={() => setVisble(false)}>
            <div className="w-full h-full bg-slate-900/75 absolute z-10 items-center justify-center flex">
                <div className="rounded-2xl py-4 border-solid border-stone-300/0.5 bg-slate-900 border-2">
                    <div className=" flex flex-row">
                        <div className=""></div>
                        <div className=""></div>
                        <div className=""></div>
                    </div>
                </div>
            </div>
        </div>
        , document.body
    )
}