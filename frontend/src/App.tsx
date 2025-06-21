import { BrowserRouter, Route, Routes } from "react-router";
import Layout from "./layout/layout";
import Home from "./pages/home";
import ProductDetail from "./pages/productDetail";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route element={<Layout />}> 
            <Route path="/" element={<Home />} />
            <Route path="/productDetail" element={<ProductDetail />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App;
