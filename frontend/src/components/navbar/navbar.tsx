import { Box, Flex, Link as ChakraLink, Text } from '@chakra-ui/react'
import React, { useState } from 'react'
import { Link } from 'react-router'
import { IoMdHome, IoMdCreate } from "react-icons/io";
import { IoReceipt, IoCartSharp, IoLogOutOutline, IoLogInOutline } from "react-icons/io5";

function Navbar() {
    const [login, setLogin] = useState(true)

    const loginHandler = () => {
        if(!login) {
            setLogin(true)
        } else {
            setLogin(false)
        }
    }

  return (
    <>
        {login ? <Box background="#ffffff" w="100%" h='70px' borderWidth='1px' shadow="14px 14px 26px 10px rgba(227,227,227,0.64)">
            <Flex gap='6' padding="15px" flexWrap="wrap" alignItems='center'>
                <Text textStyle="xl" flex='1'>BalPay</Text>
                <Text display='inline-flex' alignItems="center"><IoMdHome />&nbsp;Home</Text>
                <Text display='inline-flex' alignItems="center"><IoCartSharp />&nbsp;Top Up</Text>
                <Text display='inline-flex' alignItems="center"><IoReceipt />&nbsp;My Transaction</Text>
                <ChakraLink display='inline-flex' alignItems="center">
                   <IoLogOutOutline /> <Link to="#" onClick={loginHandler}>Logout</Link>  
                </ChakraLink>
            </Flex>
        </Box> : <Box background="#ffffff" w="100%" h='70px' borderWidth='1px' shadow="14px 14px 26px 10px rgba(227,227,227,0.64)">
            <Flex gap='4' padding="15px" flexWrap="wrap">
                <Text textStyle="xl" flex='1'>BalPay</Text>
                <ChakraLink display='inline-flex' alignItems="center">
                    <IoLogInOutline /><Link to="#" onClick={loginHandler}>Login</Link>  
                </ChakraLink>
                <Text display='inline-flex' alignItems="center"><IoMdCreate />&nbsp;Register</Text>
            </Flex>
        </Box>}
    </>
  )
}

export default Navbar