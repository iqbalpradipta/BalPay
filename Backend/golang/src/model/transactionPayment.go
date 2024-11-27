package model

import "time"

type TransactionPayment struct {
	Id                int             	`json:"id" gorm:"primaryKey:autoIncrement"`
  	PaymentStatus     string			`json:"paymentStatus"`		
  	PaidAt            time.Time         `json:"paidAt"`
  	CreatedAt         time.Time         `json:"createdAt"`
  	UpdatedAt         time.Time         `json:"updatedAt"`
  	TransactionId     int				`json:"transactionId"`
  	PaymentMethodId   int				`json:"paymentMethodId"`
  	Transaction       *Transaction      `gorm:"foreignKey:transactionId"`
  	PaymentMethod     *PaymentMethod    `gorm:"foreignKey:paymentMethodId"`
}