import type { CardProductProps } from "@/interface/ICardProduct";
import type { IProduct } from "@/interface/IProduct";
import { Box, Flex, Heading, Image, Text } from "@chakra-ui/react";
import { CiCircleInfo } from "react-icons/ci";

function DescriptionCard({name, image}: IProduct) {
  return (
    <>
      <Box p={5} bg="#f4f6fb" borderRadius="md" boxShadow="md" mb={4}>
        <Box display="flex" alignItems="center" mb={4}>
          <Image
            src={image}
            alt={name}
            boxSize="80px"
            mr={3}
            borderRadius="lg"
          />
          <Heading size="md">{name}</Heading>
        </Box>
        <Box>
          <Text fontWeight='bold' fontSize='xl'>Cara Top Up</Text> 
          <Flex gap={2} alignItems={"center"}>
            <CiCircleInfo style={{color: 'green'}} />
            Masukkan User ID kamu
          </Flex>
          <Flex gap={2} alignItems={"center"}>
            <CiCircleInfo style={{color: 'green'}}/>
            Pilih Nominal yang kamu inginkan{" "}
          </Flex>
          <Flex gap={2} alignItems={"center"}>
            <CiCircleInfo style={{color: 'green'}}/>
            Selesaikan pembayaran
          </Flex>
          <Flex gap={2} alignItems={"center"}>
            <CiCircleInfo style={{color: 'green'}}/> 
            Top up akan ditambahkan ke akun Game kamu{" "}
          </Flex>
        </Box>
      </Box>
    </>
  );
}

export default DescriptionCard;
