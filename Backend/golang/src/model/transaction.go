package model

import "time"

type Transaction struct {
	Id                  int						`json:"id" gorm:"primaryKey:autoIncrement"`
	Status              string					`json:"status" form:"status"`
	Amount              float64					`json:"amount" form:"amount"`
	CreatedAt           time.Time           	`json:"createdAt"`
	UpdatedAt           time.Time           	`json:"updatedAt"`
	UserId              int						`json:"userId"`
	GameId              int						`json:"gameId"`
	GameProductId       int						`json:"gameProductId"`
	Users               *Users            		`gorm:"foreignKey:userId"`
	Games               *Games              	`gorm:"foreignKey:gameId"`
	GameProduct         *GameProduct         	`gorm:"foreignKey:gameProductId"`
	TransactionPayment  []TransactionPayment	`gorm:"foreignKey:transactionId"`
}