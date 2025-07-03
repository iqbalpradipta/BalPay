import { Box, Flex, Image } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { Link } from "react-router";
import { IoMdHome, IoMdCreate } from "react-icons/io";
import {
  IoReceipt,
  IoCartSharp,
  IoLogOutOutline,
  IoLogInOutline,
} from "react-icons/io5";
import { toaster } from "../ui/toaster";

function Navbar() {
  const [login, setLogin] = useState(!!localStorage.getItem("authToken"));

  const handleLogout = () => {
    localStorage.removeItem("authToken");
    setLogin(false)
    toaster.create({
      title: "Logout Berhasil!",
      description: `Silahkan Login kembali untuk melanjutkan`,
      type: "success",
      closable: true,
      duration: 5000,
    });
  };

  return (
    <>
      {login ? (
        <Box
          background="#ffffff"
          shadow="14px 14px 26px 10px rgba(227,227,227,0.64)"
          pe={3}
        >
          <Flex gap="6" flexWrap="wrap" alignItems="center">
            <Box flex="1" asChild>
              <Link to="/">
                <Image
                  src="/assets/logo.png"
                  alt="logo Balpay"
                  w="60px"
                  ms={5}
                />
              </Link>
            </Box>
            <Box>
              <Link
                to="/"
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoMdHome />
                &nbsp;Home
              </Link>
            </Box>
            <Box>
              <Link
                to="/topup"
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoCartSharp />
                &nbsp;Top Up
              </Link>
            </Box>
            <Box>
              <Link
                to="/transaction"
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoReceipt />
                &nbsp;My Transaction
              </Link>
            </Box>
            <Box>
              <Link
                to="#"
                onClick={handleLogout}
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoLogOutOutline />
                &nbsp;Logout
              </Link>
            </Box>
          </Flex>
        </Box>
      ) : (
        <Box
          background="#ffffff"
          shadow="14px 14px 26px 10px rgba(227,227,227,0.64)"
          pe={3}
        >
          <Flex gap="4" flexWrap="wrap" alignItems="center">
            <Box flex="1">
              <Link to="/">
                <Image src="/assets/logo.png" w="60px" ms={5} />
              </Link>
            </Box>
            <Box>
              <Link
                to="/login"
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoLogInOutline />
                &nbsp;Login
              </Link>
            </Box>
            <Box>
              <Link
                to="/register"
                style={{ display: "inline-flex", alignItems: "center" }}
              >
                <IoMdCreate />
                &nbsp;Register
              </Link>
            </Box>
          </Flex>
        </Box>
      )}
    </>
  );
}

export default Navbar;
