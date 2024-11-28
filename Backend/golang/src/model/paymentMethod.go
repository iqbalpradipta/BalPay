package model

import "time"

type PaymentMethod struct {
	Id                  int             		`json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name                string					`json:"name" form:"name" gorm:"column:name"`
	Description         string					`json:"description" form:"description" gorm:"column:description"`
	CreatedAt   		time.Time     			`json:"createdAt" gorm:"column:createdAt"`
    UpdatedAt   		time.Time     			`json:"updatedAt" gorm:"column:updatedAt"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:paymentMethodId"`
}

func(PaymentMethod) TableName() string {
	return "PaymentMethod"
}