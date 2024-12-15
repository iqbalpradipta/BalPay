import {
    Divider,
} from "antd";
import FormProduct from "../component/formProduct";

function TopUpProduct() {
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
            <FormProduct />
            <Divider>Tentang Produk</Divider>
            <div style={{ padding: "18px", lineHeight: "1.8" }}>
                {formatDescription.map((text, index) => (
                    <p key={index}>{text.trim()}</p>
                ))}
            </div>
        </>
    )
}

export default TopUpProduct