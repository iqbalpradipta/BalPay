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
import type { IProduct } from "@/interface/IProduct";
import { toaster } from "../ui/toaster";
import api from "@/api/api";
import { useNavigate } from "react-router";

function ValueCard({ ProductDetail }: IProduct) {
  const [price, setPrice] = useState("");
  const [gameUserId, setGameUserId] = useState("");
  const [productDetailID, setProductDetailID] = useState("");
  const [purchaseQuantity, setPurchaseQuantity] = useState<number | "">(1);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const items = (ProductDetail ?? []).map((item) => ({
    ...item,
    value: item.ID.toString(),
    label: item.name,
  }));

  const productDetails = createListCollection({ items });
  const dataProduct = { gameUserId, productDetailID, purchaseQuantity };

  const handleSubmit = async () => {
    setLoading(true);
    try {
      const response = await api.post("/transaction", dataProduct);
      setGameUserId("");
      setProductDetailID("");
      setPurchaseQuantity(1);
      toaster.create({
        title: "Transaksi Success",
        description: "Transaksi Berhasil dibuat",
        type: "success",
        closable: true,
        duration: 5000,
      });
      window.open(
        `http://localhost:8000/api/v1/transaction/${response.data.data.ID}/pay`,
        "_blank"
      );
      navigate(`/invoice/${response.data.data.ID}`);
    } catch (error) {
      console.log(error);
      toaster.create({
        title: "Invalid Input",
        description: error,
        type: "error",
        closable: true,
        duration: 5000,
      });
    } finally {
      setLoading(false);
    }
  };

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

      setProductDetailID(selectedItem.ID);
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
            <Input
              placeholder="Harga"
              value={price}
              readOnly
              onChange={handleChange}
            />
          </Select.Root>

          <Heading size="md" color="blue.500" mb={2}>
            2 Masukkan jumlah pembelian
          </Heading>
          <Input
            placeholder="1"
            mb={6}
            value={purchaseQuantity}
            onChange={(e) => {
              const val = e.target.value;
              setPurchaseQuantity(val === "" ? "" : parseInt(val));
            }}
          />

          <Heading size="md" color="blue.500" mb={2}>
            3 Masukkan Detail Account
          </Heading>
          <Input
            placeholder="contoh: (uid)/(server) [*Khusus untuk pembelian steam wallet, masukkan email disini]"
            mb={6}
            value={gameUserId}
            onChange={(e) => setGameUserId(e.target.value)}
          />
        </Box>
        <Center>
          <Button
            colorPalette="blue"
            borderRadius="2xl"
            width="40%"
            loading={loading}
            loadingText={"Transaksi Dibuat ..."}
            onClick={handleSubmit}
          >
            <FaShoppingCart />
            Lanjutkan Pembayaran
          </Button>
        </Center>
      </Box>
    </>
  );
}

export default ValueCard;
