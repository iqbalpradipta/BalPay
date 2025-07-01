import ListTransaction from "@/components/transactionComponents/listTransaction";
import { Container } from "@chakra-ui/react";

function Transaction() {
  return (
    <>
      <Container fluid p={5}>
        <ListTransaction />
      </Container>
    </>
  );
}

export default Transaction;
