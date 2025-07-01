import { BrowserRouter, Route, Routes } from "react-router";
import Layout from "./layout/layout";
import Home from "./pages/home";
import ProductDetail from "./pages/productDetail";
import Invoice from "./pages/invoice";
import Transaction from "./pages/transaction";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route element={<Layout />}> 
            <Route path="/" element={<Home />} />
            <Route path="/productDetail" element={<ProductDetail />} />
            <Route path="/invoice" element={<Invoice />} />
            <Route path="/transaction" element={<Transaction />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App;
