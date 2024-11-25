export interface IGameProductData {
  games: { connect: { id: number } }
  name: string;
  price: number;
  image: string[];
  type: string;
}
