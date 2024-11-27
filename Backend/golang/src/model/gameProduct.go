package model

import "time"

type GameProduct struct {
	Id          int         	`json:"id" gorm:"primaryKey:autoIncrement"`
  	Name        string			`json:"name"`
  	Price       float64			`json:"price"`
  	Image       []string		`json:"image" gorm:"type:text[]"`
  	Type        string			`json:"type"`
  	CreatedAt   time.Time      	`json:"createdAt"`
  	UpdatedAt   time.Time      	`json:"updatedAt"`
  	GameId      int				
  	Games       *Games 			`gorm:"foreignKey:gameId"`        
  	Transaction []Transaction	`gorm:"foreignKey:gameProductId"`
}