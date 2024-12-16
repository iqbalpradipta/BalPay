import { BrowserRouter, Route, Routes } from "react-router-dom"
import LayoutComponent from "./layout"
import Home from "./pages/home"
import Product from "./pages/product"
import Order from "./pages/order"
import Login from "./pages/login"
import Register from "./pages/register"
import ProductType from "./pages/productType"
import TopUpProduct from "./pages/topUpProduct"
import AkunProduct from "./pages/akunProduct"
import DetailProduct from "./pages/detailProduct"

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LayoutComponent />}>
            <Route path="/" element={<Home />} />
            <Route path="/product" element={<Product />} />
            <Route path="/product/productType" element={<ProductType />} />
            <Route path="/product/productType/topUpProduct" element={<TopUpProduct />} />
            <Route path="/product/productType/akunProduct" element={<AkunProduct />} />
            <Route path="/product/productType/akunProduct/detailProduct" element={<DetailProduct />} />
            <Route path="/order" element={<Order />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
