export class CoinDetail {
  public id: string;
  public symbol: string | undefined;
  public name: string | undefined;
  public rank: number | undefined;
  public price: number | undefined;
  public image: string | undefined;
  public priceChangePercentage24h: number | undefined;
  public high24h: number | undefined;
  public low24h: number | undefined;


  constructor(
    id: string,
    symbol?: string,
    name?: string,
    rank?: number,
    price?: number,
    image?: string,
    priceChangePercentage24h?: number,
    high24h?: number,
    low24h?: number,
  ) {
    this.id = id;
    this.symbol = symbol;
    this.name = name;
    this.rank = rank;
    this.price = price;
    this.image = image;
    this.priceChangePercentage24h = priceChangePercentage24h;
    this.high24h = high24h;
    this.low24h = low24h;
  }
}
