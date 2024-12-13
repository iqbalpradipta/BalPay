import { Descriptions, DescriptionsProps, Divider, Typography } from 'antd'

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
            span: 'filled',
            children: "Diamond 100"
        }
    ]

    return (
        <>  
            <Typography.Title level={4} style={{flex: 1, textAlign: 'center'}}>Order Status</Typography.Title>
            <Descriptions bordered items={items} style={{padding: 20}} />
            <Divider />
            <Descriptions bordered items={items} style={{padding: 20}} />
            <Divider />
            <Descriptions bordered items={items} style={{padding: 20}} />
            <Divider />
        </>
    )
}

export default Order