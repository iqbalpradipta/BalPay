package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name		string `json:"name" form:"name" gorm:"not null"`
	Description string `json:"description" form:"description"`
	Image		string `json:"image" form:"image"`

	ProductDetail []ProductDetail `gorm:"foreignKey:ProductID"`
}