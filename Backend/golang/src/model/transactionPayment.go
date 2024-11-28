package model

import "time"

type TransactionPayment struct {
	Id              int           `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	PaymentStatus   string        `json:"paymentStatus" form:"paymentStatus" gorm:"column:paymentStatus"`
	PaidAt          time.Time     `json:"paidAt" gorm:"column:paidAt"`
	CreatedAt       time.Time     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt" gorm:"column:updatedAt"`
	TransactionId   int           `json:"transactionId" gorm:"column:transactionId"`
	PaymentMethodId int           `json:"paymentMethodId" gorm:"column:paymentMethodId"`
	Transaction     Transaction   `gorm:"foreignKey:transactionId"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:paymentMethodId"`
}

func (TransactionPayment) TableName() string {
	return "TransactionPayment"
}
