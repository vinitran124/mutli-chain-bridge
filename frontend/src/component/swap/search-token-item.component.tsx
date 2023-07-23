import { FaEthereum } from 'react-icons/fa'
import { BsCheck2 } from 'react-icons/bs'

export const SearchTokenItem = () => {
    return (<div className="flex flex-row p-3 items-center justify-between">
        <div className='flex flex-row items-center'>
            <FaEthereum color='white' width={48} height={48} scale={1.5} className="mr-3" />
            <div>
                <div className="text-white font-medium text-base">Ether</div>
                <div className="text-gray-400 font-thin text-xs">ETH</div>
            </div>
        </div>
        <BsCheck2 color='#4c82fb' />
    </div>)
}