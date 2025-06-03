package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionCode		string `json:"transactionCode" form:"transactionCode" gorm:"unique"`
	GameUserId			string `json:"gameUserId" form:"gameUserId"`
	StatusTransaction 	string `json:"statusTransaction" form:"statusTransaction"`

	UserID				int 
	PaymentMethodID 	int
	ProductDetailID		int

	User				User
	PaymentMethod 		PaymentMethod
	ProductDetail		ProductDetail
}