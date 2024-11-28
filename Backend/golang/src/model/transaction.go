package model

import "time"

type Transaction struct {
	Id                  int						`json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Status              string					`json:"status" form:"status" gorm:"column:status"`
	Amount              float64					`json:"amount" form:"amount" gorm:"column:amount"`
	CreatedAt   		time.Time     			`json:"createdAt" gorm:"column:createdAt"`
    UpdatedAt   		time.Time     			`json:"updatedAt" gorm:"column:updatedAt"`
	UserId              int						`json:"userId" form:"userId" gorm:"column:userId"`
	GameId              int						`json:"gameId" form:"gameId" gorm:"column:gameId"`
	GameProductId       int						`json:"gameProductId" form:"gameProductId" gorm:"column:gameProductId"`
	Users               Users            		`gorm:"foreignKey:userId"`
	Games               Games              		`gorm:"foreignKey:gameId"`
	GameProduct         GameProduct         	`gorm:"foreignKey:gameProductId"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:transactionId"`
}

func(Transaction) TableName() string {
	return "Transaction"
}