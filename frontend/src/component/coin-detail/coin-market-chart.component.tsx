import {
  XAxis,
  AreaChart,
  YAxis,
  CartesianGrid,
  Tooltip,
  Area,
} from 'recharts';
import { CoinMarketChartModel } from '../../model/market-chart.model';

export const CoinMarketChart = ({
  datas,
}: {
  datas: CoinMarketChartModel[];
}) => {
  return (
    <AreaChart
      width={730}
      height={250}
      data={datas}
      margin={{ top: 10, right: 30, left: 0, bottom: 0 }}
    >
      <defs>
        <linearGradient id="colorPrice" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stopColor="#704CE3" stopOpacity={1} />
          <stop offset="95%" stopColor="#534CF7" stopOpacity={0.4} />
        </linearGradient>
      </defs>
      <XAxis dataKey="date" />
      <YAxis
        type="number"
        domain={[
          Math.max(...datas.map(item => item.price)),
          Math.min(...datas.map(item => item.price)),
        ]}
      />
      <CartesianGrid strokeDasharray="3 3" />
      <Tooltip content={CustomTooltip}/>
      <Area
        type="monotone"
        dataKey="price"
        stroke="#534CF7"
        dot={false}
        fillOpacity={1}
        fill="url(#colorPrice)"
      />
    </AreaChart>
  );
};

const CustomTooltip = ({ active, payload, label }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="bg-white text-[red]">
        <p className="">{`$${payload[0].payload?.price}`}</p>
        <p className="">{`${payload[0]?.payload?.date}`}</p>
      </div>
    );
  }

  return null;
};