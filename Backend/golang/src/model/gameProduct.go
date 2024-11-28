package model

import "time"

type GameProduct struct {
	Id          int         	`json:"id" gorm:"primaryKey;autoIncrement;column:id"`
  	Name        string			`json:"name" form:"name" gorm:"column:name"`
  	Price       float64			`json:"price" form:"price" gorm:"column:price"`
  	Image       string			`json:"image" form:"image" gorm:"column:image"`
  	Type        string			`json:"type" form:"type" gorm:"column:type"`
  	CreatedAt   time.Time     	`json:"createdAt" gorm:"column:createdAt"`
    UpdatedAt   time.Time     	`json:"updatedAt" gorm:"column:updatedAt"`
  	GameId      int				`json:"gameId" gorm:"column:gameId"`
  	Games       Games 			`gorm:"foreignKey:gameId"`        
  	Transaction []Transaction	`gorm:"foreignKey:gameProductId"`
}

func(GameProduct) TableName() string {
	return "GameProduct"
}