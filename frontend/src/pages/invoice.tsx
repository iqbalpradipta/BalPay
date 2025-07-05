import api from "@/api/api";
import { toaster } from "@/components/ui/toaster";
import type { TransactionDetail } from "@/interface/ITransaction";
import {
  Box,
  Heading,
  Stack,
  Flex,
  Icon,
  Text,
  Center,
  Button,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoReceipt } from "react-icons/io5";
import { LuScrollText } from "react-icons/lu";
import { MdLightbulb } from "react-icons/md";
import { useParams } from "react-router";
import { Link } from "react-router";

function Invoice() {
  const { id } = useParams();
  const [transaction, setTransaction] = useState<TransactionDetail | null>();
  const format = (price: any) => {
    return Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
    }).format(price);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await api.get(
          `http://localhost:8000/api/v1/transaction/${id}`
        );
        setTransaction(res.data.data);
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

  return (
    <>
      <Center w="100%" h="calc(100vh)">
        <Box
          bg="white"
          borderRadius="md"
          p={5}
          boxShadow="sm"
          w="full"
          maxW="lg"
          fontSize="sm"
        >
          <Box>
            <Center>
              <LuScrollText size="150px" color="blue" />
            </Center>
            <Center>
              <Text fontSize="md" fontWeight="bold">
                Detail Transaksi
              </Text>
            </Center>
          </Box>

          <Heading size="sm" mb={4} mt={5}>
            Rincian Pembayaran
          </Heading>

          <Stack gap={2}>
            <Flex justify="space-between" fontWeight="semibold">
              <Text color="gray.600">Transaction Code</Text>
              <Text fontWeight="semibold">{transaction?.transactionCode}</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">Game Account ID</Text>
              <Text>187904827298249</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">Status Pembayaran</Text>
              <Text>{transaction?.statusTransaction}</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">Nama Produk</Text>
              <Text>
                {transaction?.ProductDetail.name}{" "}
                <span>{`(x ${transaction?.purchaseQuantity} Item)`}</span>
              </Text>
            </Flex>
          </Stack>

          <Box height="1px" bg="gray.200" my={4} w="100%" />

          <Flex justify="space-between" align="center">
            <Text fontWeight="semibold">Total Belanja</Text>
            <Text fontSize="lg" fontWeight="bold">
              {format(transaction?.totalTransaction)}
            </Text>
          </Flex>

          <Flex align="center" mt={2} color="gray.500" fontSize="xs" gap={1}>
            <Icon as={MdLightbulb} boxSize={4} color="green.400" />
            <Text>Lakukan Pembayaran Segera</Text>
          </Flex>
          <Center mt={3}>
            <Button asChild colorPalette="blue" borderRadius="2xl" width="100%">
              <Link to="/transaction">
                <IoReceipt />
                Cek Transaction
              </Link>
            </Button>
          </Center>
        </Box>
      </Center>
    </>
  );
}

export default Invoice;
