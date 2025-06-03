package model

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name		string			`json:"name" form:"name"`

	Transaction	[]Transaction	`gorm:"foreignKey:PaymentMethodID"`
}