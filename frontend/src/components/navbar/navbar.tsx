import { Box, Flex, Text } from '@chakra-ui/react'
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
                <Box textStyle="xl" flex='1'>
                    <Link to="/">Bal<span style={{color: 'blue', fontWeight: 'bold'}}>Pay</span></Link>
                </Box>
                <Box>
                    <Link to='/' style={{display: 'inline-flex', alignItems: 'center'}}><IoMdHome />&nbsp;Home</Link>
                </Box>
                <Box>
                    <Link to='/topup' style={{display: 'inline-flex', alignItems: 'center'}}><IoCartSharp />&nbsp;Top Up</Link>
                </Box>
                <Box>
                    <Link to='/transaction' style={{display: 'inline-flex', alignItems: 'center'}}><IoReceipt />&nbsp;My Transaction</Link>
                </Box>
                <Box>
                   <IoLogOutOutline style={{display: 'inline-flex', alignItems: 'center'}} />&nbsp;<Link to="#" onClick={loginHandler}>Logout</Link>  
                </Box>
            </Flex>
        </Box> : <Box background="#ffffff" w="100%" h='70px' borderWidth='1px' shadow="14px 14px 26px 10px rgba(227,227,227,0.64)">
            <Flex gap='4' padding="15px" flexWrap="wrap">
                <Box textStyle="xl" flex='1'>
                    <Link to="/">BalPay</Link>
                </Box>
                <Box>
                    <IoLogInOutline style={{display: 'inline-flex', alignItems: 'center'}} />&nbsp;<Link to="#" onClick={loginHandler}>Login</Link>  
                </Box>
                <Box>
                    <Link to='/register' style={{display: 'inline-flex', alignItems: 'center'}}><IoMdCreate />&nbsp;Register</Link>
                </Box>
            </Flex>
        </Box>}
    </>
  )
}

export default Navbar