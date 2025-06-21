import Navbar from "@/components/navbar/navbar";
import { Box } from "@chakra-ui/react";
import { Outlet } from "react-router";

function Layout() {
  return (
    <>
      <Box width='100%' height='100%' bgColor='#f4f6fb'>
          <Navbar />
          <Outlet />
      </Box>
    </>
  );
}

export default Layout;
