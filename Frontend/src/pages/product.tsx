import {
    Col,
    Row,
    Image,
    Button,
    Form,
    Input,
    Select,
    Divider,
} from "antd";
import productML from '../mocks/productML.json';
import { useState } from "react";

function Product() {
    const [price, setPrice] = useState<number>(0);

    const handleChangeProduct = (value: string) => {
        const product = productML.find((data) => data.name === value);
        if (product) {
            setPrice(Number(product.price));
        }
    };

    const description = `
            BalPay menawarkan top up Mobile Legends yang mudah, aman, dan instan.
            Pembayaran tersedia melalui pulsa (Telkomsel, Indosat, Tri, XL, Smartfren), Codacash, QRIS, GoPay, OVO, DANA, ShopeePay, LinkAja, Krevido, Alfamart, Indomaret, DOKU, Bank Transfer, dan Card Payments.
            Top up ML Diamond, Twilight Pass, Weekly Diamond Pass (WDP), dan Starlight hanya dalam hitungan detik! 
            Cukup masukkan User ID dan Zone ID MLBB Anda, pilih jumlah Diamond yang Anda inginkan, selesaikan pembayaran, 
            dan Diamond akan secara langsung ditambahkan ke akun Mobile Legends Anda.
            Pelajari lebih lanjut cara untuk membeli Starlight Membership & Coupon Pass, di sini.
            Login ke akun BalPay Anda dan dapatkan akses ke promo Mobile Legends dan event lainnya. Belum punya akun BalPay? Daftar sekarang!
    `
    const formatDescription = description.split(/(?<=\.)\s+/)

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
                            <Form.Item label="User ID" required>
                                <Input placeholder="Masukkan User ID Anda" />
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
            <Divider>Tentang Produk</Divider>
            <div style={{ padding: "18px", lineHeight: "1.8" }}>
                    {formatDescription.map((text, index) => (
                        <p key={index}>{text.trim()}</p>
                    ))}
            </div>
        </>
    );
}

export default Product;