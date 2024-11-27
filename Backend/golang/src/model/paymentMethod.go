package model

import "time"

type PaymentMethod struct {
	Id                  int             		`json:"id" gorm:"primaryKey:autoIncrement"`
	Name                string					`json:"name"`
	Description         string					`json:"description"`
	CreatedAt           time.Time        		`json:"createdAt"`
	UpdatedAt           time.Time        		`json:"updatedAt"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:paymentMethodId"`
}