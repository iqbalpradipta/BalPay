package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionCode   string `json:"transactionCode" form:"transactionCode" gorm:"unique"`
	GameUserId        string `json:"gameUserId" form:"gameUserId"`
	StatusTransaction string `json:"statusTransaction" form:"statusTransaction"`

	UserID          uint `json:"userId" form:"userId"`
	ProductDetailID uint `json:"productDetailId" form:"productDetailId"`

	User           User           `json:"-" gorm:"foreignKey:UserID"`
	ProductDetail  ProductDetail  `json:"-" gorm:"foreignKey:ProductDetailID"`

	Payment        Payment        `json:"-" gorm:"foreignKey:TransactionID"`
}

