export interface Product {
  name: string;
  image: string;
  description: string;
}

export interface CardProductProps {
  name?: string;
  products: Product[];
}