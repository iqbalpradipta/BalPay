import { Box } from "@chakra-ui/react";

function ValueCard() {
  return (
    <>
      <Box p={5} bg="#f4f6fb" borderRadius="md" boxShadow="md" mb={4}>
        <Box display="flex" alignItems="center" mb={4}>
          FORM VALUE
        </Box>
      </Box>
    </>
  );
}

export default ValueCard;
