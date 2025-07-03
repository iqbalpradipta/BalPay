import { toaster } from "@/components/ui/toaster";
import { UserSchema } from "@/validation/userSchema";
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
} from "@chakra-ui/react";
import axios from "axios";
import { useState, type SyntheticEvent } from "react";
import { FaGoogle } from "react-icons/fa";
import { Link, useNavigate } from "react-router";

function RegisterPage() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [errors, setErrors] = useState<{
    name?: string;
    email?: string;
    password?: string;
  } | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: SyntheticEvent) => {
    e.preventDefault();
    setLoading(true);
    setErrors(null);

    const userData = { name, email, password };

    try {
      await UserSchema.validate(userData, { abortEarly: false });
      const response = await axios.post(
        "http://localhost:8000/api/v1/user",
        userData
      );

      toaster.create({
        title: "Register telah berhasil!",
        description: "Silahkan Login untuk melanjutkan.",
        type: "success",
        closable: true,
        duration: 5000,
      });

      setName("");
      setEmail("");
      setPassword("");
      navigate("/login");
    } catch (error: any) {
      if (error.status == 500) {
        toaster.create({
          title: "Registrasi Gagal.",
          description: `Email ${email} telah digunakan. Silahkan gunakan email Aktif yang lain.`,
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
      setLoading(false);
    }
  };

  return (
    <>
      <Box border="1px" bgColor="#f4f6fb" p={7}>
        <Center>
          <Box bgColor="white" borderRadius="lg" boxShadow="lg" w="50%">
            <Flex direction="column" alignItems="center" gap={2} mt={2}>
              <Image src="/assets/logo.png" alt="logo Balpay" w="10%" />
              <Text fontWeight="bold" fontSize="xl">
                Selamat Datang
              </Text>
              <Text fontWeight="normal" fontSize="md">
                Silahkan Mendaftar Terlebih Dahulu!
              </Text>
              <Separator size="sm" mt={5} w="95%" />
              <Box p={5} w="400px">
                <Field.Root gap={2}>
                  <Field.Label>Full Name</Field.Label>
                  <Input
                    placeholder="Full Name"
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                  />
                  <Field.Label>Email</Field.Label>
                  <Input
                    placeholder="your@email.com"
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                  />
                  <Field.ErrorText>Your email is wrong!</Field.ErrorText>
                  <Field.Label>Password</Field.Label>
                  <Input
                    placeholder="yourpassword"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                  />
                </Field.Root>
                <Button
                  w="100%"
                  mt={5}
                  colorPalette="blue"
                  variant="solid"
                  onClick={handleSubmit}
                  loading={loading}
                  loadingText={"Sedang Mendaftar ...."}
                >
                  Register
                </Button>
                <Button w="100%" mt={3} colorPalette="blue" variant="outline">
                  <FaGoogle /> Register with Google
                </Button>
                <Text
                  textAlign="center"
                  mt={3}
                  textStyle="sm"
                  fontWeight="semibold"
                >
                  You already Have Account?{" "}
                  <Link to="/login" style={{ color: "#6182f8" }}>
                    Login Here!
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

export default RegisterPage;
