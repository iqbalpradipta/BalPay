
export interface IProductDetail {
  ID: string;
  name: string;
  price: number;
  Product: IProduct;
}

export interface IProduct {
  ID?: string;
  name?: string;
  image?: string;
  description?: string;
  ProductDetail?: IProductDetail[]
}

export interface IProductTransaksi {
    Product?: IProduct
    ProductDetail?: IProductDetail
}

