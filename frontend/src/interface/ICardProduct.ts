export interface Product {
  title: string;
  image: string;
  discount: string;
  label: string;
  sold: string;
}

export interface CardProductProps {
  title?: string;
  products: Product[];
}