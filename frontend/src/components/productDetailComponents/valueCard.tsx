import {
  Box,
  Heading,
  Text,
  Button,
  Input,
  Center,
  Portal,
  Select,
  createListCollection,
} from "@chakra-ui/react";
import { FaShoppingCart } from "react-icons/fa";
import { useState } from "react";

const ProductList = createListCollection({
  items: [
    { label: "60 Genesis Crystal", value: "60", price: "Rp. 12.100" },
    { label: "330 Genesis Crystal", value: "330", price: "Rp. 61.300" },
    { label: "1090 Genesis Crystal", value: "1090", price: "Rp. 184.200" },
    { label: "2240 Genesis Crystal", value: "2240", price: "Rp. 399.500" },
  ],
});

function ValueCard() {
  const [price, setPrice] = useState("");

  const handleChange = (details: any) => {
    const selectedItem = details.items;
    setPrice(selectedItem[0].price);
  };

  return (
    <>
      <Box p={5} bg="white" borderRadius="md" boxShadow="md" mb={4}>
        <Box>
          <Heading size="md" color="blue.500" mb={2}>
            1 Pilih nominal
          </Heading>
          <Select.Root
            collection={ProductList}
            size="md"
            mb={4}
            onValueChange={handleChange}
          >
            <Select.HiddenSelect />
            <Select.Control>
              <Select.Trigger>
                <Select.ValueText placeholder="Select Product" />
              </Select.Trigger>
              <Select.IndicatorGroup>
                <Select.Indicator />
              </Select.IndicatorGroup>
            </Select.Control>
            <Portal>
              <Select.Positioner>
                <Select.Content>
                  {ProductList.items.map((product) => (
                    <Select.Item item={product} key={product.value}>
                      {product.label}
                      <Select.ItemIndicator />
                    </Select.Item>
                  ))}
                </Select.Content>
              </Select.Positioner>
            </Portal>
            <Text fontSize="sm" mt={2} fontWeight="semibold" color="blue.300">
              Price
            </Text>
            <Input placeholder="Harga" value={price} readOnly />
          </Select.Root>

          <Heading size="md" color="blue.500" mb={2}>
            2 Masukkan jumlah pembelian
          </Heading>
          <Input placeholder="1" defaultValue="1" mb={6} />

          {/* Section 3 */}
          <Heading size="md" color="blue.500" mb={2}>
            3 Masukkan Detail Account
          </Heading>
          <Input placeholder="contoh: (uid)/(server)" mb={6} />
        </Box>
        <Center>
          <Button colorPalette="blue" w="50%" borderRadius="2xl" boxShadow="xl">
            <FaShoppingCart />
            Bayar
          </Button>
        </Center>
      </Box>
    </>
  );
}

export default ValueCard;
