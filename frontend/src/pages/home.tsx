import CardProduct from "@/components/homeComponents/cardProduct";
import { Box } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import api from "@/api/api";
import type { Product } from "@/interface/ICardProduct";


function Home() {
  const [data, setData] = useState<Product[]>([])

  useEffect(() => {
    const fetchData = async() => {
      try {
        const response = await api.get("/product")
        setData(response.data.data)
      } catch (error) {
        return error
      }
    }
    fetchData()
  }, [data])
  
  return (
    <>
      <Box p={1} m={3} >
        <CardProduct products={data} />
      </Box>
    </>
  );
}

export default Home;
