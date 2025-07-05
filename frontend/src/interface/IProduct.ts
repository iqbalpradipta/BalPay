export interface IProductDetail {
  ID: string;
  name: string;
  price: number
}

export interface IProduct {
  ID?: string;
  name?: string;
  image?: string;
  description?: string;
  ProductDetail?: IProductDetail[]
}
