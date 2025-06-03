package model

import "gorm.io/gorm"

type ProductDetail struct {
	gorm.Model
	Name  string  `json:"name" form:"name"`
	Stock int     `json:"stock" form:"stock"`
	Price float64 `json:"price" form:"price"`

	ProductID int

	Product Product

	Transaction	[]Transaction `gorm:"foreignKey:ProductDetailID"`
}