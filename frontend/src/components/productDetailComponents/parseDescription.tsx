import { Heading, Text } from "@chakra-ui/react";
import { List } from "@chakra-ui/react";
import type { JSX } from "react";
import { CiCircleInfo } from "react-icons/ci";


export const ParseDescription = (desc: string): JSX.Element[] => {
  const lines = desc.split("\n").filter((line) => line.trim() !== "");

  const result: JSX.Element[] = [];
  let isInList = false;
  let listItems: string[] = [];

  const flushList = () => {
    if (listItems.length > 0) {
      result.push(
        <List.Root key={`list-${Math.random()}`} gap={1} mb={4} variant="plain">
          {listItems.map((item, idx) => (
            <List.Item key={idx}>
              <List.Indicator asChild color="green.500">
                <CiCircleInfo />
              </List.Indicator>
              {item}
            </List.Item>
          ))}
        </List.Root>
      );
      listItems = [];
    }
  };

  lines.forEach((line, index) => {
    if (/^\d+\./.test(line)) {
      isInList = true;
      listItems.push(line.replace(/^\d+\.\s*/, ""));
    } else {
      if (isInList) {
        flushList();
        isInList = false;
      }

      if (
        line === "Cara Top Up" ||
        line.startsWith("Top Up") ||
        line.startsWith("Beli ")
      ) {
        result.push(
          <Heading size="sm" mb={1} mt={4} key={`heading-${index}`}>
            {line}
          </Heading>
        );
      } else {
        result.push(
          <Text fontSize="sm" mb={2} key={`text-${index}`}>
            {line}
          </Text>
        );
      }
    }
  });

  flushList();

  return result;
};
