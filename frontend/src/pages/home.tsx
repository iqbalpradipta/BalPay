import CardProduct from "@/components/homeComponents/cardProduct";
import { Box } from "@chakra-ui/react";
import product from "@/mocks/product.json"


function Home() {
  return (
    <>
      <Box p={1} m={3}>
        <CardProduct products={product} title="List Game" />
      </Box>
    </>
  );
}

export default Home;
