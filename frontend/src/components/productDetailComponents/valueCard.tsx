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
import { Link } from "react-router";
import type { IProduct } from "@/interface/IProduct";

function ValueCard({ ProductDetail }: IProduct) {
  const items = (ProductDetail ?? []).map((item) => ({
    ...item,
    value: item.ID.toString(),
    label: item.name,
  }));

  const productDetails = createListCollection({ items });
  const [price, setPrice] = useState("");

  const handleChange = (details: any) => {
    const selectedItem = items.find(
      (item) => item.value === String(details.value)
    );
    if (selectedItem) {
      setPrice(
        Intl.NumberFormat("id-ID", {
          style: "currency",
          currency: "IDR",
          minimumFractionDigits: 0,
        }).format(selectedItem.price)
      );
    }
  };

  return (
    <>
      <Box p={5} bg="white" borderRadius="md" boxShadow="md" mb={4}>
        <Box>
          <Heading size="md" color="blue.500" mb={2}>
            1 Pilih nominal
          </Heading>
          <Select.Root
            collection={productDetails}
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
                  {items?.map((product) => (
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
          <Input placeholder="1" defaultValue={1} mb={6} />

          <Heading size="md" color="blue.500" mb={2}>
            3 Masukkan Detail Account
          </Heading>
          <Input
            placeholder="contoh: (uid)/(server) [*Khusus untuk pembelian steam wallet, masukkan email disini]"
            mb={6}
          />
        </Box>
        <Center>
          <Button asChild colorPalette="blue" borderRadius="2xl" width="40%">
            <Link to="/invoice">
              <FaShoppingCart />
              Lanjutkan Pembayaran
            </Link>
          </Button>
        </Center>
      </Box>
    </>
  );
}

export default ValueCard;
