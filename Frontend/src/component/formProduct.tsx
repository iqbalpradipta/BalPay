import {
    Col,
    Row,
    Image,
    Button,
    Form,
    Input,
    Select,
} from "antd";
import productML from '../mocks/productML.json';
import { useState } from "react";

function FormProduct() {
    const [price, setPrice] = useState<number>(0);

    const handleChangeProduct = (value: string) => {
        const product = productML.find((data) => data.name === value);
        if (product) {
            setPrice(Number(product.price));
        }
    };
    return (
        <>
            <div style={{ padding: 30 }}>
                <Row gutter={[16, 16]}>
                    <Col span={12}>
                        <Image
                            src="https://news.codashop.com/id/wp-content/uploads/sites/4/2021/01/ID-MLBB-Lucky-draw-Januari21-banner.jpg"
                            alt="Mobile Legends Banner"
                            preview={false}
                        />
                    </Col>
                    <Col span={12}>
                        <Form
                            labelCol={{ span: 6 }}
                            wrapperCol={{ span: 16 }}
                            layout="horizontal"
                            style={{ maxWidth: 800 }}
                        >
                            <Form.Item label="User ID & Zone ID" required>
                                <Input placeholder="contoh: 90001238042(2018)" />
                            </Form.Item>
                            <Form.Item label="Product" required>
                                <Select
                                    placeholder="Pilih Produk"
                                    onChange={handleChangeProduct}
                                >
                                    {productML.map((data) => (
                                        <Select.Option key={data.name} value={data.name}>
                                            {data.name}
                                        </Select.Option>
                                    ))}
                                </Select>
                            </Form.Item>
                            <Form.Item label="Price">
                                <Input
                                    value={price}
                                    readOnly
                                    placeholder="Harga akan muncul di sini"
                                />
                            </Form.Item>
                            <Form.Item>
                                <Button
                                    type="primary"
                                    style={{
                                        marginLeft: '38%',
                                        width: '60%',
                                        fontWeight: 'bold',
                                        backgroundColor: "#001529",
                                    }}
                                >
                                    Beli Sekarang!
                                </Button>
                            </Form.Item>
                        </Form>
                    </Col>
                </Row>
            </div>
        </>
    )
}

export default FormProduct