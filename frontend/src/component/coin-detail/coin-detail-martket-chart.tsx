import axios from 'axios';
import { ISeriesApi, Time, createChart } from 'lightweight-charts';
import { Interval } from '../../const/interval.const';
import { useEffect } from 'react';

export const CoinDetailMarketChart = ({
  symbol,
  interval = Interval._1m,
  limit = 200,
}: {
  symbol: string | undefined;
  interval?: Interval;
  limit?: number;
}) => {
  let candleSeries: ISeriesApi<'Candlestick', Time> | undefined = undefined;
  let timerMarketChart: NodeJS.Timer;

  // only use in unmount component
  useEffect(() => {
    return () => {
      if (timerMarketChart) {
        clearInterval(timerMarketChart);
      }
    };
  }, []);

  
  useEffect(() => {
    if (timerMarketChart) clearInterval(timerMarketChart);

    timerMarketChart = setInterval(() => {
      syncMarketChart(false);
    }, 1000);
  }, [symbol, interval, limit]);

  const initChart = () => {
    const domElement = document.getElementById('coin-detail-market-chart');
    if (
      symbol &&
      candleSeries === undefined &&
      domElement &&
      !domElement.hasChildNodes()
    ) {
      console.log('into1');
      const chartProperties = {
        layout: {
          background: {
            // type: 'solid',
            color: '#0F172A',
          },
          // lineColor: '#2B2B43',
          textColor: '#D9D9D9',
        },
        watermark: {
          color: 'rgba(0, 0, 0, 0)',
        },
        crosshair: {
          // color: '#758696',
        },
        grid: {
          vertLines: {
            color: '#2B2B43',
          },
          horzLines: {
            color: '#363C4E',
          },
        },
        timeScale: {
          timeVisible: true,
          secondsVisible: false,
        },
      };
      const chart = createChart(domElement, chartProperties);
      candleSeries = chart.addCandlestickSeries();
      // syncMarketChart(false)

      timerMarketChart = setInterval(() => {
        syncMarketChart(false);
      }, 1000);
    }
  };

  const syncMarketChart = async (isUpdateChart: boolean) => {
    console.log(`syncMarketChart: http://127.0.0.1:9665/fetchAPI?endpoint=https://api.binance.com/api/v3/klines?symbol=${symbol?.toUpperCase()}USDT&interval=${interval}&limit=${limit}`)
    if (candleSeries) {
      await axios(
        `http://127.0.0.1:9665/fetchAPI?endpoint=https://api.binance.com/api/v3/klines?symbol=${symbol?.toUpperCase()}USDT&interval=${interval}&limit=${limit}`,
      )
        .then((res: any) => {
          const cdata = res.data.map((d: any) => {
            return {
              time: d[0] / 1000,
              open: parseFloat(d[1]),
              high: parseFloat(d[2]),
              low: parseFloat(d[3]),
              close: parseFloat(d[4]),
            };
          });
          if (candleSeries) {
            if (isUpdateChart) {
              candleSeries.update(cdata);
            } else {
              candleSeries.setData(cdata);
            }
            console.log('cdata:', cdata);
          }
        })
        .catch(err => console.log(err));
    }
  };

  initChart();

  return (
    <div className="flex flex-row w-full h-full">
      {interval}
      <div id="coin-detail-market-chart" className="w-full h-full"></div>
    </div>
  );
};
