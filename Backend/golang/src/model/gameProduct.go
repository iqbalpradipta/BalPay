package model

import "time"

type GameProduct struct {
	Id          int         	`json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
  	Name        string			`json:"name,omitempty" form:"name" gorm:"column:name"`
  	Price       float64			`json:"price,omitempty" form:"price" gorm:"column:price"`
  	Image       string			`json:"image,omitempty" form:"image" gorm:"column:image"`
  	Type        string			`json:"type,omitempty" form:"type" gorm:"column:type"`
  	CreatedAt   time.Time     	`json:"createdAt,omitempty" gorm:"column:createdAt"`
    UpdatedAt   time.Time     	`json:"updatedAt,omitempty" gorm:"column:updatedAt"`
  	GameId      int				`json:"gameId,omitempty" gorm:"column:gameId"`
  	Games       Games 			`gorm:"foreignKey:gameId" json:"Games,omitempty"`        
  	Transaction []Transaction	`gorm:"foreignKey:gameProductId" json:"Transaction,omitempty"`
}

func(GameProduct) TableName() string {
	return "GameProduct"
}