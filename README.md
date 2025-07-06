
# BalPay — Website Pembelian Top Up Game dan Voucher Digital Terintegrasi Payment Gateway (Xendit)
BalPay adalah aplikasi website untuk pembelian top-up game, voucher digital, dan layanan pembayaran online, dilengkapi sistem transaksi otomatis terhubung ke Xendit Payment Gateway.

## ✨ Features

- ✅ **Autentikasi Pengguna**
  - Login/Register dengan Email & Password
  - Integrasi Google OAuth2

- ✅ **Pembelian Produk Game Digital**
  - Produk dengan harga dan deskripsi
  - Form pemesanan dengan input validasi

- ✅ **Pembayaran Otomatis dengan Xendit**
  - Invoice dibuat otomatis
  - Redirect ke halaman pembayaran Xendit
  - Status transaksi diperbarui otomatis melalui webhook

- ✅ **Dashboard Transaksi**
  - Riwayat transaksi pengguna
  - Status: pending / paid
  - Detail transaksi lengkap


## 📸 Tampilan Antarmuka
| Login/Register | Home (List Game) |
|---|---|
| ![Login](https://res.cloudinary.com/dszok6ewm/image/upload/v1751757776/ProjectImage/BalPay/Register_kcxi3f.jpg) | ![Home](https://res.cloudinary.com/dszok6ewm/image/upload/v1751757777/ProjectImage/BalPay/Home_om17z2.jpg) |

| Detail Produk | Daftar Transaksi |
|---|---|
| ![Detail Produk](https://res.cloudinary.com/dszok6ewm/image/upload/v1751758080/ProjectImage/BalPay/DetailProduct_uw7bxl.jpg) | ![Transaksi](https://res.cloudinary.com/dszok6ewm/image/upload/v1751757776/ProjectImage/BalPay/Transaction_esk23x.jpg) |

| Invoice |
|---|
| ![Invoice](https://res.cloudinary.com/dszok6ewm/image/upload/v1751757777/ProjectImage/BalPay/Invoice_luwm9r.jpg) |

## 🧩 Teknologi yang Digunakan

### 🔹 Frontend
- `React 19`
- `React Router v7`
- `Chakra UI v3.21.0`
- `Redux Toolkit`
- `Axios`
- `Yup` (validasi)
- `React Icons`
- `next-themes` (opsional dark mode)

### 🔹 Backend (Golang)
- `Echo v4`
- `GORM + MySQL`
- `JWT Authentication (v5)`
- `Xendit Go SDK`
- `godotenv` (env config)
**Integrasi Cloudinary tersedia** via `cloudinary-go` untuk upload gambar jika dibutuhkan.

## 📂 Struktur Penting Backend

- `/model` – Struct model (User, Product, Transaction, Payment)
- `/repository` – Layer akses data (GORM)
- `/services` – Logika bisnis: CreateInvoice, HandleWebhook, dsb
- `/handler` – HTTP handler Echo untuk endpoint
- `/routes` – Routing utama
- `/middleware` – Auth middleware
- `.env` – XENDIT_API_KEY, DB, JWT_SECRET

## 🔗 Integrasi Xendit
- Create invoice via `xendit-go/invoice`
- Simpan `payment_token`, `payment_url`, `external_id`
- Webhook otomatis dari Xendit ke endpoint:  
  `POST /api/v1/webhook/xendit`
- Validasi `status`, update `payment.status` dan `transaction.status_transaction
## 📂 API Documentation
You can see API Spesification with this link:

**Swagger BalPay**: (https://app.swaggerhub.com/apis-docs/IQBALPRADIPTA2/BalPay/1.0.0)

![Login](https://res.cloudinary.com/dszok6ewm/image/upload/v1751764079/ProjectImage/BalPay/API_Documentation_kg0e4j.jpg)
## 🛠️ Instalasi & Menjalankan Project

### 🔸 Frontend
```bash
cd frontend
npm install
npm run dev
```
### 🔷 Backend
```bash
cd backend
go mod tidy
go run main.go
```

### ⚠ .env Setting
```bash
DB_USER=root
DB_PASS=yourpass
DB_HOST=localhost
DB_PORT=yourPort
DB_NAME=(name.db) eg: db_balpay
SECRETKEY=BEBAS
CLOUDINARY_URL=your_cloudinary_key
XENDIT_API_KEY=your_xendit_key
```
    
## Authors

- [Iqbal Pradipta](https://www.github.com/iqbalpradipta)

