import { BrowserRouter, Route, Routes } from "react-router-dom"
import LayoutComponent from "./layout"
import Home from "./pages/home"

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LayoutComponent />}>
            <Route path="/" element={<Home />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
