import { CoinDetail } from '../../model/coin-detail.model';
import Slider from '@mui/material/Slider';

export const CoinInformation = ({ data }: { data: CoinDetail }) => {
  return (
    <div className="flex flex-col">
      <div className="flex flex-row items-center space-x-[5px]">
        <div className="w-[40px] h-[40px]">
          <img src={data.image} alt="--" />
        </div>
        <div className="text-[16px] font-bold">{data.name}</div>
        <div className="text-[14px]">- {data.symbol?.toUpperCase()}</div>
        <div className="text-[12px] text-gray-400">#{data.rank}</div>
      </div>

      <div className="flex flex-row items-center space-x-1">
        <div className="text-[36px] font-bold">${data.price}</div>
        <div
          className={
            'text-[18px] font-bold ' +
            ((data.priceChangePercentage24h ?? 0) < 0
              ? 'text-red-500'
              : 'text-green-500')
          }
        >
          {(data.priceChangePercentage24h ?? 0) < 0
            ? `${data.priceChangePercentage24h}`
            : `+${data.priceChangePercentage24h}`}
          %
        </div>
      </div>

      <div className="flex flex-col">
        <div className='mx-[8px]'>
          <Slider
            size="small"
            disabled
            aria-label="Small"
            defaultValue={
              ((data.price ?? 0) - (data.low24h ?? 0)) /
              ((data.high24h ?? 0) - (data.low24h ?? 0) * 100)
            }
          />
        </div>
        <div className="flex flex-row justify-between text-[12px] font-[600]">
          <div>${data.low24h}</div>
          <div>Price in 24h</div>
          <div>${data.high24h}</div>
        </div>
      </div>
    </div>
  );
};
