import { Box, Center, Flex, Text, Image, Field, Input, Separator, Checkbox, Grid, GridItem, Button } from "@chakra-ui/react";
import { FaGoogle } from "react-icons/fa"
import { Link } from "react-router";

function LoginPage() {
  return (
    <>
      <Box border="1px" bgColor="#f4f6fb" p={7}>
        <Center>
          <Box bgColor='white' borderRadius='lg' boxShadow='lg' w='50%'>
            <Flex direction="column" alignItems="center" gap={2} mt={1}>
              <Image src="/assets/logo.png" alt="logo Balpay" w="10%" />
              <Text fontWeight='bold' fontSize='xl'>Selamat Datang</Text>
              <Text fontWeight='normal' fontSize='md'>Silahkan login untuk melanjutkan!</Text>
              <Separator size='sm' mt={5} w='95%'  />
                <Box p={5} w='400px'>
                    <Field.Root gap={2}>
                      <Field.Label>Email</Field.Label>
                      <Input placeholder="your@email.com" type='email' />
                      <Field.ErrorText>Your email is wrong!</Field.ErrorText>
                      <Field.Label>Password</Field.Label>
                      <Input placeholder="yourpassword" type='password'/>
                    </Field.Root>
                    <Grid templateColumns="repeat(2, 1fr)">
                      <GridItem>
                        <Checkbox.Root mt={4} value="remember me">
                          <Checkbox.HiddenInput />
                          <Checkbox.Control />
                          <Checkbox.Label>Remember me</Checkbox.Label>
                        </Checkbox.Root>
                      </GridItem>
                      <GridItem >
                        <Text mt={4} textAlign='end' textStyle='sm' color='blue.500'>
                          <Link to='forgot-password'>Forgot Password?</Link>
                        </Text>
                      </GridItem>
                    </Grid>
                    <Button w='100%' mt={5} colorPalette='blue' variant='solid'>Login</Button>
                    <Button w='100%' mt={3} colorPalette='blue' variant='outline'><FaGoogle /> Login with Google</Button>
                    <Text textAlign='center' mt={3} textStyle='sm' fontWeight='semibold'>Don't have an account? <Link to='/register' style={{color: '#6182f8'}}>Register Here!</Link></Text>
                </Box>
            </Flex>
          </Box>
        </Center>
      </Box>
    </>
  );
}

export default LoginPage;
