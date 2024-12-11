import React from 'react'
import Navbar from './component/navbar'
import { Outlet } from 'react-router-dom'
import { Layout } from 'antd'

function LayoutComponent() {
  return (
    <>
      <Layout>
        <Navbar />
        <Outlet />
      </Layout>
    </>
  )
}

export default LayoutComponent