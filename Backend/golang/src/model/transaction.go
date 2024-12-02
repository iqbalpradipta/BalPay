package model

import "time"

type Transaction struct {
	Id                  int						`json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	Status              string					`json:"status,omitempty" form:"status" gorm:"column:status"`
	Amount              float64					`json:"amount,omitempty" form:"amount" gorm:"column:amount"`
	CreatedAt   		time.Time     			`json:"createdAt,omitempty" gorm:"column:createdAt"`
    UpdatedAt   		time.Time     			`json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	UserId              int						`json:"userId,omitempty" form:"userId" gorm:"column:userId"`
	GameId              int						`json:"gameId,omitempty" form:"gameId" gorm:"column:gameId"`
	GameProductId       int						`json:"gameProductId,omitempty" form:"gameProductId" gorm:"column:gameProductId"`
	Users               Users            		`gorm:"foreignKey:userId" json:"Users,omitempty"`
	Games               Games              		`gorm:"foreignKey:gameId" json:"Games,omitempty"`
	GameProduct         GameProduct         	`gorm:"foreignKey:gameProductId" json:"GameProduct,omitempty"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:transactionId" json:"TransactionPayment,omitempty"`
}

func(Transaction) TableName() string {
	return "Transaction"
}