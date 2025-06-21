import {
  Box,
  Card,
  Heading,
  Image,
  Text,
  Badge,
  SimpleGrid,
  HStack,
  Icon,
  Button,
} from "@chakra-ui/react";
import { FaShoppingCart } from "react-icons/fa";
import type { CardProductProps } from "@/interface/ICardProduct";
import { Link } from "react-router";

const CardProduct = ({ title, products }: CardProductProps) => {
  return (
    <>
      <Box bg={"#f4f6fb"} p={5} borderRadius="md">
        {title && (
          <Heading fontSize="xl" mb={4}>
            {title}
          </Heading>
        )}
        <SimpleGrid columns={[2, 3, 6]} gap={4}>
          {products.map((game, index) => (
            <Card.Root
              key={index}
              borderRadius="lg"
              overflow="hidden"
              boxShadow="sm"
            >
              <Image src={game.image} alt={game.title} h="100%" w="100%" />
              <Card.Body>
                <Text fontWeight="semibold">{game.title}</Text>
                <HStack mt={2} gap={2}>
                  <Badge colorScheme="red">{game.discount}</Badge>
                  <Badge colorScheme="blue" variant="subtle">
                    {game.label}
                  </Badge>
                </HStack>
              </Card.Body>
              <Card.Footer>
                <Button asChild colorPalette='cyan'>
                  <Link to='/productDetail'><FaShoppingCart />Beli Sekarang</Link>
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
