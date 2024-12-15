import { Button, Descriptions, DescriptionsProps, Divider, Typography } from 'antd'
import { Link } from 'react-router-dom'

function Order() {
    const items: DescriptionsProps['items'] = [
        {
            label: "Order ID",
            span: 'filled',
            children: "20"
        },
        {
            label: "Status Order",
            span: 'filled',
            children: "Pending"
        },
        {
            label: "Name Product",
            span: 2,
            children: "Diamond 100"
        }
    ]

    return (
        <>
            <Typography.Title level={4} style={{ flex: 1, textAlign: 'center' }}>Order Status</Typography.Title>
            <Descriptions bordered extra={
                <Button variant='outlined' color='default' style={{ marginTop: 0 }}>
                    <Link to='/product/productType'>Bayar Sekarang</Link>
                </Button>
            }
                items={items}
                style={{ padding: 20 }}
            />
            <Divider />
            <Descriptions bordered extra={
                <Button variant='outlined' color='default' style={{ marginTop: 0 }}>
                    <Link to='/product/productType'>Bayar Sekarang</Link>
                </Button>
            }
                items={items}
                style={{ padding: 20 }}
            />
            <Divider />
            <Descriptions bordered extra={
                <Button variant='outlined' color='default' style={{ marginTop: 0 }}>
                    <Link to='/product/productType'>Bayar Sekarang</Link>
                </Button>
            }
                items={items}
                style={{ padding: 20 }}
            />
            <Divider />
        </>
    )
}

export default Order