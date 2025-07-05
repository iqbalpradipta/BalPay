export interface Product {
  ID: string;
  name: string;
  image: string;
  description: string;
}

export interface CardProductProps {
  name?: string;
  products: Product[];
}