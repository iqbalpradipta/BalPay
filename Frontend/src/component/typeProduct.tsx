import { Row, Col, Card, Button } from 'antd'
import { IoGameController } from 'react-icons/io5';
import { GrMoney } from "react-icons/gr";
import { Link } from 'react-router-dom';

const { Meta } = Card;

function TypeProduct() {
  return (
    <>
      <div style={{ marginTop: 100, marginBottom: 155 }}>
        <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 32 }}>
          <Col span={12} className='gutter-row'>
            <Card hoverable style={{ width: 300, padding: 20 }} cover={<IoGameController style={{ width: 200, height: 100 }} />}>
              <Meta title="Beli Akun" description="Temukan akun yang sesuai kebutuhanmu disini !" />
              <Button variant='outlined' color='default' style={{ marginTop: 20 }}>
                <Link to='/product/productType/akunProduct'>Beli Sekarang!</Link>
              </Button>
            </Card>
          </Col>
          <Col span={12}>
            <Card hoverable style={{ width: 300, padding: 20 }} cover={<GrMoney style={{ width: 200, height: 100 }} />}>
              <Meta title="Top Up Akun" description="Top Up Anti Ribet sesuai keinginanmu disini !" />
              <Button variant='outlined' color='default' style={{ marginTop: 20 }}>
                <Link to='/product/productType/topUpProduct'>Beli Sekarang!</Link>
              </Button>
            </Card>
          </Col>
        </Row>
      </div>
    </>
  )
}

export default TypeProduct