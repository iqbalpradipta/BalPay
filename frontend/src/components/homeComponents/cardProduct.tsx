import {
  Box,
  Card,
  Heading,
  Image,
  Text,
  SimpleGrid,
  HStack,
  Button,
} from "@chakra-ui/react";
import { FaShoppingCart } from "react-icons/fa";
import type { CardProductProps } from "@/interface/ICardProduct";
import { Link, useNavigate } from "react-router";

const CardProduct = ({ products }: CardProductProps) => {
  const navigate = useNavigate()
  return (
    <>
      <Box bg={"#f4f6fb"} p={5} borderRadius="md" boxShadow="lg">
        <Heading fontSize="xl" mb={4}>
          List Game
        </Heading>
        <SimpleGrid columns={[2, 3, 6]} gap={4}>
          {products.map((game, index) => (
            <Card.Root
              key={index}
              borderRadius="lg"
              overflow="hidden"
              boxShadow="sm"
            >
              <Image src={game.image} alt={game.name} h="100%" w="100%" />
              <Card.Body>
                <Text fontWeight="semibold">{game.name}</Text>
                <HStack mt={2} gap={2}>
                  <Text>{game.description}</Text>
                </HStack>
              </Card.Body>
              <Card.Footer>
                <Button
                  asChild
                  colorPalette="blue"
                  boxShadow="lg"
                  borderRadius="lg"
                >
                  <Link to={`/productDetail/${game.ID}`}>
                    <FaShoppingCart />
                    Beli Sekarang
                  </Link>
                </Button>
              </Card.Footer>
            </Card.Root>
          ))}
        </SimpleGrid>
      </Box>
    </>
  );
};

export default CardProduct;
