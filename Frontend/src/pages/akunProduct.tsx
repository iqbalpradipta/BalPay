import { Card, Button, Typography, Row, Col } from 'antd'
import { Link } from 'react-router-dom'
import akunML from '../mocks/akunML.json'

const { Meta } = Card;

function AkunProduct() {
    return (
        <>
            <div style={{ padding: 10, marginBottom: 190, marginTop: 10 }}>
                <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 32 }}>
                    {akunML.map((data, index) => (
                        <Col className="gutter-row" key={index} span={6} xs={24} sm={12} md={8} lg={6}>
                            <Card
                                style={{ width: 300, marginTop: 10 }}
                                cover={
                                    <img
                                        alt={data.name}
                                        src={data.image}
                                        style={{width: 300, height: 220}}
                                    />
                                }
                                actions={[
                                    <Button variant='outlined' color='default'>
                                        <Link to='/product/productType/akunProduct/detailProduct'>Lihat Detail</Link>
                                    </Button>,
                                    <Button variant='outlined' color='default'>
                                        <Link to='/product/productType/akunProduct'>Beli Sekarang!</Link>
                                    </Button>
                                ]}
                            >
                                <Meta
                                    title={data.name}
                                    description={data.deskripsi}
                                />
                                <Typography.Title level={5} style={{ marginTop: 10, position: 'relative', top: 20 }}>Rp. {data.price}</Typography.Title>
                            </Card>
                        </Col>
                    ))}
                </Row>
            </div>
        </>
    )
}

export default AkunProduct