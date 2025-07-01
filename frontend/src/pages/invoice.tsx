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
import { IoReceipt } from "react-icons/io5";
import { LuScrollText } from "react-icons/lu";
import { MdLightbulb } from "react-icons/md";
import { Link } from "react-router";

function Invoice() {
  return (
    <>
      <Center w="100%" h="calc(88vh)">
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
            <Flex justify="space-between">
              <Text color="gray.600">Metode Pembayaran</Text>
              <Text fontWeight="semibold">GoPay</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">Status Pembayaran</Text>
              <Text>Unpaid</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">330 Genesis Crystal</Text>
              <Text>Rp61.300</Text>
            </Flex>

            <Flex justify="space-between">
              <Text color="gray.600">Pajak</Text>
              <Text>Rp1.000</Text>
            </Flex>
          </Stack>

          <Box height="1px" bg="gray.200" my={4} w="100%" />

          <Flex justify="space-between" align="center">
            <Text fontWeight="semibold">Total Belanja</Text>
            <Text fontSize="lg" fontWeight="bold">
              Rp158.500
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
