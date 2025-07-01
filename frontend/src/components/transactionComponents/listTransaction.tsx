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

import barang from "@/mocks/barang.json";
import { HiChevronLeft, HiChevronRight } from "react-icons/hi";
import { useState, useMemo } from "react";

const pageSize = 4;

function ListTransaction() {
  const [page, setPage] = useState(1);

  const totalItems = barang.length;
  const totalPages = Math.ceil(totalItems / pageSize);

  const visibleItems = useMemo(() => {
    const startRange = (page - 1) * pageSize;
    const endRange = startRange + pageSize;
    return barang.slice(startRange, endRange);
  }, [page, barang, pageSize]);

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
              </Table.ColumnGroup>
              <Table.Header>
                <Table.Row>
                  <Table.ColumnHeader> </Table.ColumnHeader>
                  <Table.ColumnHeader>Product Name</Table.ColumnHeader>
                  <Table.ColumnHeader>Status</Table.ColumnHeader>
                  <Table.ColumnHeader textAlign="end">Price</Table.ColumnHeader>
                </Table.Row>
              </Table.Header>
              <Table.Body>
                {visibleItems.map((item) => (
                  <Table.Row key={item.id}>
                    <Table.Cell>
                      <Image
                        src={item.image}
                        alt={item.game}
                        width="80px"
                        p={2}
                        ms={5}
                      />
                    </Table.Cell>
                    <Table.Cell>
                      <Flex direction="column">
                        <Text fontSize="md" fontWeight="bold">
                          {item.game}
                        </Text>
                        <Text fontSize="small">{item.label}</Text>
                      </Flex>
                    </Table.Cell>
                    <Table.Cell>{item.status}</Table.Cell>
                    <Table.Cell textAlign="end">{item.price}</Table.Cell>
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
                <Pagination.Item
                  type="page"
                  value={pageNumber}
                  asChild={true}
                >
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
