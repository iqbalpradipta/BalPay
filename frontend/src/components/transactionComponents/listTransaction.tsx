import {
  Box,
  Flex,
  Image,
  Text,
  Stack,
  ButtonGroup,
  IconButton,
  Table,
  Button,
} from "@chakra-ui/react";
import { Pagination } from "@ark-ui/react/pagination";
import { HiChevronLeft, HiChevronRight } from "react-icons/hi";
import { useState, useMemo, useEffect } from "react";
import api from "@/api/api";
import { toaster } from "../ui/toaster";
import type { TransactionDetail } from "@/interface/ITransaction";
import { Link } from "react-router";
import { IoReceipt } from "react-icons/io5";

const pageSize = 4;

function ListTransaction() {
  const [page, setPage] = useState(1);
  const [data, setData] = useState<TransactionDetail[]>([]);

  const totalItems = data.length;
  const totalPages = Math.ceil(totalItems / pageSize);
  
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
        const response = await api.get(`/transaction`);
        setData(response.data.data);
      } catch (error) {
        toaster.create({
          title: "Invalid Input",
          description: error,
          type: "error",
          closable: true,
          duration: 5000,
        });
      }
    };
    fetchData();
  }, []);

  const visibleItems = useMemo(() => {
    const startRange = (page - 1) * pageSize;
    const endRange = startRange + pageSize;
    return data.slice(startRange, endRange);
  }, [page, data, pageSize]);

  const pageNumbers = Array.from({ length: totalPages }, (_, i) => i + 1);

  return (
    <>
      <Box bg="white" borderRadius="md" p={5} boxShadow="sm" fontSize="sm">
        <Stack gap="4">
          <Stack>
            <Table.Root size="sm" variant="outline" borderRadius="md">
              <Table.ColumnGroup>
                <Table.Column htmlWidth="10%" />
                <Table.Column />
                <Table.Column />
                <Table.Column />
                <Table.Column />
                <Table.Column />
              </Table.ColumnGroup>
              <Table.Header>
                <Table.Row>
                  <Table.ColumnHeader></Table.ColumnHeader>
                  <Table.ColumnHeader>Product Name</Table.ColumnHeader>
                  <Table.ColumnHeader>Invoice Code</Table.ColumnHeader>
                  <Table.ColumnHeader>Status</Table.ColumnHeader>
                  <Table.ColumnHeader>Price</Table.ColumnHeader>
                  <Table.ColumnHeader textAlign="end"></Table.ColumnHeader>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {visibleItems.map((item) => (
                  <Table.Row key={item.ID}>
                    <Table.Cell>
                      <Image
                        src={item.ProductDetail.Product?.image}
                        alt={item.ProductDetail.Product?.name}
                        width="80px"
                        p={2}
                        ms={5}
                      />
                    </Table.Cell>
                    <Table.Cell>
                      <Flex direction="column">
                        <Text fontSize="md" fontWeight="bold">
                          {item.ProductDetail.Product?.name}
                        </Text>
                        <Text fontSize="small">{format(item.totalTransaction)}</Text>
                      </Flex>
                    </Table.Cell>
                    <Table.Cell>{item.transactionCode}</Table.Cell>
                    <Table.Cell>{item.statusTransaction}</Table.Cell>
                    <Table.Cell>{format(item.totalTransaction)}</Table.Cell>
                    <Table.Cell textAlign="end">
                      <Button asChild colorPalette="blue" borderRadius="2xl">
                        <Link to={`/invoice/${item.ID}`}>
                          <IoReceipt />
                          Cek Transaction
                        </Link>
                      </Button>
                    </Table.Cell>
                  </Table.Row>
                ))}
              </Table.Body>
            </Table.Root>
          </Stack>

          <Pagination.Root
            count={totalItems}
            pageSize={pageSize}
            page={page}
            onPageChange={(details) => setPage(details.page)}
          >
            <ButtonGroup variant="ghost" size="sm">
              <Pagination.PrevTrigger asChild>
                <IconButton aria-label="Previous Page">
                  <HiChevronLeft />
                </IconButton>
              </Pagination.PrevTrigger>

              {pageNumbers.map((pageNumber) => (
                <Pagination.Item type="page" value={pageNumber} asChild={true}>
                  <Button variant={pageNumber === page ? "outline" : "ghost"}>
                    {pageNumber}
                  </Button>
                </Pagination.Item>
              ))}

              <Pagination.NextTrigger asChild>
                <IconButton aria-label="Next Page">
                  <HiChevronRight />
                </IconButton>
              </Pagination.NextTrigger>
            </ButtonGroup>
          </Pagination.Root>
        </Stack>
      </Box>
    </>
  );
}

export default ListTransaction;
