import { toaster } from "@/components/ui/toaster";
import type { RootState } from "@/redux/store";
import { login, setEmail, setPassword } from "@/redux/userSlice";
import { LoginSchema } from "@/validation/userSchema";
import {
  Box,
  Center,
  Flex,
  Text,
  Image,
  Field,
  Input,
  Separator,
  Button,
  Checkbox,
  Grid,
  GridItem,
} from "@chakra-ui/react";
import axios from "axios";
import { type SyntheticEvent } from "react";
import { FaGoogle } from "react-icons/fa";
import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router";

function LoginPage({
  setLogin,
}: {
  setLogin: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const user = useSelector((state: RootState) => state.user);
  const dispatch = useDispatch();
  const userData = { email: user.email, password: user.password };

  const navigate = useNavigate();

  const handleLogin = async (e: SyntheticEvent) => {
    e.preventDefault();
    dispatch(login(true));

    try {
      await LoginSchema.validate(userData, { abortEarly: false });
      const response = await axios.post(
        "http://localhost:8000/api/v1/login",
        userData
      );

      const token = response.data.data;
      localStorage.setItem("authToken", token);

      setLogin(true);
      dispatch(setEmail(""))
      dispatch(setPassword(""))  
      navigate("/");
      toaster.create({
        title: "Login Berhasil!",
        description: `Selamat Datang ${user.email} :)`,
        type: "success",
        closable: true,
        duration: 5000,
      });

    } catch (error: any) {
      if (error.status == 500) {
        toaster.create({
          title: "Login Gagal.",
          description: `Username atau password salah.`,
          type: "error",
          closable: true,
          duration: 8000,
        });
      } else {
        toaster.create({
          title: "Invalid Input",
          description: error.message,
          type: "error",
          closable: true,
          duration: 5000,
        });
      }
    } finally {
      dispatch(login(false));
    }
  };
  return (
    <>
      <Box border="1px" bgColor="#f4f6fb" p={7}>
        <Center>
          <Box bgColor="white" borderRadius="lg" boxShadow="lg" w="50%">
            <Flex direction="column" alignItems="center" gap={2} mt={1}>
              <Image src="/assets/logo.png" alt="logo Balpay" w="10%" />
              <Text fontWeight="bold" fontSize="xl">
                Selamat Datang
              </Text>
              <Text fontWeight="normal" fontSize="md">
                Silahkan login untuk melanjutkan!
              </Text>
              <Separator size="sm" mt={5} w="95%" />
              <Box p={5} w="400px">
                <Field.Root gap={2}>
                  <Field.Label>Email</Field.Label>
                  <Input
                    placeholder="your@email.com"
                    type="email"
                    value={user.email}
                    onChange={(e) => dispatch(setEmail(e.target.value))}
                  />
                  <Field.ErrorText>Your email is wrong!</Field.ErrorText>
                  <Field.Label>Password</Field.Label>
                  <Input
                    placeholder="yourpassword"
                    type="password"
                    value={user.password}
                    onChange={(e) => dispatch(setPassword(e.target.value))}
                    onKeyDown={(e) => {
                      if(e.key == 'Enter') {
                        handleLogin(e)
                      }
                    }}
                  />
                </Field.Root>
                <Grid templateColumns="repeat(2, 1fr)">
                  <GridItem>
                    <Checkbox.Root mt={4} value="remember me">
                      <Checkbox.HiddenInput />
                      <Checkbox.Control />
                      <Checkbox.Label>Remember me</Checkbox.Label>
                    </Checkbox.Root>
                  </GridItem>
                  <GridItem>
                    <Text
                      mt={4}
                      textAlign="end"
                      textStyle="sm"
                      color="blue.500"
                    >
                      <Link to="forgot-password">Forgot Password?</Link>
                    </Text>
                  </GridItem>
                </Grid>
                <Button
                  w="100%"
                  mt={5}
                  colorPalette="blue"
                  variant="solid"
                  onClick={handleLogin}
                  loading={user.loading}
                  loadingText={"Mencoba Login ...."}
                >
                  Login
                </Button>
                <Button w="100%" mt={3} colorPalette="blue" variant="outline">
                  <FaGoogle /> Login with Google
                </Button>
                <Text
                  textAlign="center"
                  mt={3}
                  textStyle="sm"
                  fontWeight="semibold"
                >
                  Don't have an account?{" "}
                  <Link to="/register" style={{ color: "#6182f8" }}>
                    Register Here!
                  </Link>
                </Text>
              </Box>
            </Flex>
          </Box>
        </Center>
      </Box>
    </>
  );
}

export default LoginPage;
