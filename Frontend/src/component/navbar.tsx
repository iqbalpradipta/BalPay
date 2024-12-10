import { Menu } from 'antd'
import { Header } from 'antd/es/layout/layout'
import { IoIosHome, IoIosLogIn } from "react-icons/io";
import { FaBasketShopping } from "react-icons/fa6";
import { AiFillProduct } from "react-icons/ai";
import { FaUserPlus } from "react-icons/fa";
import Search, { SearchProps } from 'antd/es/input/Search';
import { Typography } from 'antd';
import { Link } from 'react-router-dom';

function Navbar() {
  const onSearch: SearchProps['onSearch'] = (value, _e, info) => console.log(info?.source, value);
  const { Title } = Typography;
  const items = [
    {
      label: (<Link to={'/'}>Home</Link>),
      key: "home",
      icon: <IoIosHome />
    },
    {
      label: (<Link to={'/product'}>Product</Link>),
      key: "product",
      icon: <AiFillProduct />
    },
    {
      label: (<Link to={'/order'}>Order</Link>),
      key: "order",
      icon: <FaBasketShopping />
    },
    {
      label: (<Link to={'/login'}>Login</Link>),
      key: "login",
      icon: <IoIosLogIn />
    },
    {
      label: (<Link to={'/register'}>Register</Link>),
      key: "register",
      icon: <FaUserPlus />
    }
  ]

  return (
    <>
      <Header style={{ display: 'flex', gap: 50, alignItems: 'center' }}>
        <Title level={3} style={{ color: 'white', padding: 10 }} >BalPay</Title>
        <Search placeholder="input search text" onSearch={onSearch} style={{
          flex: 1,
        }} />
        <Menu
          theme="dark"
          mode="horizontal"
          items={items}
        />
      </Header>
    </>
  )
}

export default Navbar