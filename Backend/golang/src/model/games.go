package model

import "time"

type Games struct {
	Id          int          	`json:"id" gorm:"primaryKey;autoIncrement;column:id"`
  	Name        string			`json:"name" form:"name" gorm:"column:name"`
  	Description string			`json:"description" form:"description" gorm:"column:description"`
  	Icon        string			`json:"icon" form:"icon" gorm:"column:icon"`
  	Types       string			`json:"types" form:"types" gorm:"column:types"`
  	CreatedAt   time.Time     	`json:"createdAt" gorm:"column:createdAt"`
    UpdatedAt   time.Time     	`json:"updatedAt" gorm:"column:updatedAt"`
  	GameProduct []GameProduct	`gorm:"foreignKey:gameId"`
  	Transaction []Transaction	`gorm:"foreignKey:gameId"`
}

func(Games) TableName() string {
	return "Games"
}