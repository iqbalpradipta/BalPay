package model

import "time"

type TransactionPayment struct {
	Id              int           `json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	PaymentStatus   string        `json:"paymentStatus,omitempty" form:"paymentStatus" gorm:"column:paymentStatus"`
	PaidAt          time.Time     `json:"paidAt,omitempty" gorm:"column:paidAt"`
	CreatedAt       time.Time     `json:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	TransactionId   int           `json:"transactionId,omitempty" gorm:"column:transactionId"`
	PaymentMethodId int           `json:"paymentMethodId,omitempty" gorm:"column:paymentMethodId"`
	Transaction     Transaction   `gorm:"foreignKey:transactionId" json:"Transaction,omitempty"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:paymentMethodId" json:"PaymentMethod,omitempty"`
}

func (TransactionPayment) TableName() string {
	return "TransactionPayment"
}
