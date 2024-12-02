package model

import "time"

type PaymentMethod struct {
	Id                  int             		`json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	Name                string					`json:"name,omitempty" form:"name" gorm:"column:name"`
	Description         string					`json:"description,omitempty" form:"description" gorm:"column:description"`
	CreatedAt   		time.Time     			`json:"createdAt,omitempty" gorm:"column:createdAt"`
    UpdatedAt   		time.Time     			`json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:paymentMethodId" json:"TransactionPayment,omitempty"`
}

func(PaymentMethod) TableName() string {
	return "PaymentMethod"
}