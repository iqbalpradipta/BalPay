import DescriptionCard from "@/components/productDetailComponents/descriptionCard";
import ValueCard from "@/components/productDetailComponents/valueCard";
import { Container } from "@chakra-ui/react";

function ProductDetail() {
  return (
    <>
      <Container maxW='4xl' p={3}>
        <DescriptionCard />
        <ValueCard />
      </Container>
    </>
  );
}

export default ProductDetail;
