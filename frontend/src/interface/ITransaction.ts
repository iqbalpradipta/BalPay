import type { IProductDetail, IProductTransaksi } from "./IProduct"

export interface TransactionDetail {
    ID: string
    transactionCode: string
    totalTransaction: number
    statusTransaction: string
    purchaseQuantity: number
    gameUserId: string
    ProductDetail: IProductDetail
}