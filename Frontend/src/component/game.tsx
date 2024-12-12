import { Card, Col, Button, Row } from 'antd'
import game from '../mocks/game.json'
import { Link } from 'react-router-dom'

function Game() {
    const { Meta } = Card
    return (
        <>
            <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 32 }}>
                {game.map((value, index) => (
                    <Col className="gutter-row" span={6} key={index} xs={24} sm={12} md={8} lg={6}>
                        <div style={{ padding: 20 }}  >
                            <Card
                                hoverable
                                style={{ width: '80%' }}
                                cover={<img alt={value.judul} src={value.image} />}
                            >
                                <Meta title={value.judul} description={value.deskripsi} />
                                <Button variant='outlined' color='default' style={{ marginTop: 20 }}>
                                    <Link to='/product'>Beli Sekarang!</Link>
                                </Button>
                            </Card>
                        </div>
                    </Col>
                ))}
            </Row>
        </>
    )
}

export default Game