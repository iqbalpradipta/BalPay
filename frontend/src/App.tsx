import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
  Outlet,
  useNavigate,
} from "react-router";
import Layout from "./layout/layout";
import Home from "./pages/home";
import ProductDetail from "./pages/productDetail";
import Invoice from "./pages/invoice";
import Transaction from "./pages/transaction";
import LoginPage from "./pages/authentication/loginPage";
import RegisterPage from "./pages/authentication/registerPage";
import { useEffect, useState } from "react";

function App() {
  const [login, setLogin] = useState(() => !!localStorage.getItem("authToken"));

  const ProtectedRoute: React.FC<{ isAllowed: boolean }> = ({ isAllowed }) => {
    return isAllowed ? <Outlet /> : <Navigate to="/login" replace />;
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage setLogin={setLogin} />} />
        <Route path="/register" element={<RegisterPage />} />

        <Route element={<ProtectedRoute isAllowed={login} />}>
          <Route element={<Layout />}>
            <Route path="/" element={<Home />} />
            <Route path="/productDetail/:id" element={<ProductDetail />} />
            <Route path="/invoice" element={<Invoice />} />
            <Route path="/transaction" element={<Transaction />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
