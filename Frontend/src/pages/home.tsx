import { Card, Col, Divider, Button, Row } from 'antd'
import Banner from '../component/banner'
import game from '../mocks/game.json'

function Home() {

  const { Meta } = Card

  return (
    <>
      <Banner />
      <Divider style={{ border: 1, borderColor: 'gray' }}>Product</Divider>
      <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 32 }}>
        {game.map((value, index) => (
          <Col className="gutter-row" span={6}>
            <div style={{ padding: 20 }} key={index} >
              <Card
                hoverable
                style={{ width: 240 }}
                cover={<img alt={value.judul} src={value.image} />}
              >
                <Meta title={value.judul} description={value.deskripsi} />
                <Button variant='outlined' color='default'style={{marginTop: 20}}>
                  Buy Now
                </Button>
              </Card>
            </div>
          </Col>
        ))}
      </Row>
    </>
  )
}

export default Home