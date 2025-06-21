import { ParseDescription } from "@/components/productDetailComponents/parseDescription";
import { Box, Heading, Image } from "@chakra-ui/react";
const product = {
  title: "Genshin Impact",
  image:
    "https://image.api.playstation.com/vulcan/ap/rnd/202408/2010/6e7d87fef87405e9925e810a1620df04c3b98c2086711336.png",
  description: `Cara Top Up
1. Masukkan User ID dan Pilih Server kamu
2. Pilih Nominal Welkin atau Crystals yang kamu inginkan
3. Selesaikan pembayaran
4. Top up akan ditambahkan ke akun Genshin Impact kamu

Top Up Genshin Impact Termurah di Lapakgaming
Top up Genshin termurah, cepat, dan dijamin aman. Dapatkan Blessing Welkin Pack dan Crystals Genshin Impact untuk berbagai kebutuhan kamu menjelajah Teyvat cuma di Lapakgaming!

Beli Welkin Moon Dapat Diskon 30% di Lapakgaming
Dapatkan diskon 30% Welkin Moon Genshin khusus pengguna baru di Lapakgaming! Pakai kode promo AUTOKIGH saat checkout untuk dapatin diskonnya. Beli Welkin Genshin sekarang di Lapakgaming!`,
};

function DescriptionCard() {
  return (
    <>
      <Box p={5} bg="#f4f6fb" borderRadius="md" boxShadow="md" mb={4}>
        <Box display="flex" alignItems="center" mb={4}>
          <Image
            src={product.image}
            alt={product.title}
            boxSize="80px"
            mr={3}
            borderRadius="lg"
          />
          <Heading size="md">{product.title}</Heading>
        </Box>

        {ParseDescription(product.description)}
      </Box>
    </>
  );
}

export default DescriptionCard;
