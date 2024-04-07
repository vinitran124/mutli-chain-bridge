import { useState } from 'react';
import { useParams } from 'react-router-dom';

import moment from 'moment';
import axios from 'axios';
import { CoinMarketChartModel } from '../model/market-chart.model';
import { CoinMarketChart } from '../component/coin-detail/coin-market-chart.component';
import { CoinInformation } from '../component/coin-detail/coin-information.component';
import { CoinDetail } from '../model/coin-detail.model';

export interface MarketChartFilter {
  fixedPrice: number;
  currency: string;
  days: number;
}

export const CoinDetailScreen = () => {
  const { paramId } = useParams();
  const [filter, setFilter] = useState<MarketChartFilter>({
    fixedPrice: 3,
    currency: 'usd',
    days: 1,
  });
  const [marketChartData, setMarketChartData] = useState<
    CoinMarketChartModel[]
  >([]);
  const [coinDetail, setCoinDetail] = useState<CoinDetail | undefined>();

  //   const timerInterval = setInterval(() => {
  //     // getData();

  //     console.log('now: ', moment(Date.now()).format('YYYY-MM-DD HH:mm:ss'));
  //     // console.log("timerInterval", timerInterval);
  //     // console.log("data", data);
  //   }, 5000);

  const getData = async () => {
    const searchId = paramId ?? 'bella-protocol';
    console.log("paramId", paramId);
    console.log("paramId", searchId);
    // get detail
    await axios
      .get(`https://api.coingecko.com/api/v3/coins/${searchId}`)
      .then(res => {
        setCoinDetail(
          new CoinDetail(
            res.data.id,
            res.data.symbol,
            res.data.name,
            res.data.market_cap_rank,
            Number(res.data.market_data.current_price[filter.currency]),
            res.data.image.large,
            res.data.market_data.price_change_percentage_24h,
            res.data.market_data.high_24h[filter.currency],
            res.data.market_data.low_24h[filter.currency],
          ),
        );
      });

    // get martket chart
    await axios
      .get(
        `https://api.coingecko.com/api/v3/coins/${searchId}/market_chart?vs_currency=${filter.currency}&days=${filter.days}`,
      )
      .then(async res => {
        await setMarketChartData(
          res.data.prices.map((price: any) => {
            return {
              date: moment(price[0]).format('DD-MM-yyyy HH:mm:ss'),
              price: Number(price[1]).toFixed(3),
            };
          }),
        );
        console.log('res.data:', res.data);
        console.log('data:', marketChartData);
      });
  };

  return (
    <div className="pt-[80px] px-[24px] bg-white overflow-auto">
      {coinDetail ? (
        <div className='flex flex-row space-x-[24px] justify-between'>
          <div className="overflow-auto">
            <CoinInformation data={coinDetail} />
          </div>
          <div className=''>
            <CoinMarketChart datas={marketChartData} />
          </div>
        </div>
      ) : (
        <div>Coin not found</div>
      )}
      <button onClick={getData}>Sync data</button>
    </div>
  );
};
