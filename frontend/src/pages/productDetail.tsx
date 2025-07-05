import api from "@/api/api";
import DescriptionCard from "@/components/productDetailComponents/descriptionCard";
import ValueCard from "@/components/productDetailComponents/valueCard";
import { toaster } from "@/components/ui/toaster";
import type { IProduct } from "@/interface/IProduct";
import { Container } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useParams } from "react-router";

function ProductDetail() {
  const { id } = useParams();
  const [product, setProduct] = useState<IProduct | null>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await api.get(`http://localhost:8000/api/v1/product/${id}`);
        setProduct(res.data.data);
      } catch (error) {
        console.log(error);
        toaster.create({
          title: "Error Fetch Data!",
          description: `Something error when fetch data.`,
          type: "error",
          closable: true,
          duration: 5000,
        });
      }
    };
    fetchData();
  }, [id]);

  console.log(product);

  return (
    <>
      {product && (
        <Container maxW="4xl" p={3}>
          <DescriptionCard name={product.name} image={product.image} />
          <ValueCard ProductDetail={product.ProductDetail} />
        </Container>
      )}
    </>
  );
}

export default ProductDetail;
