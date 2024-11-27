package model

import "time"

type Games struct {
	Id          int          	`json:"id" gorm:"primaryKey:autoIncrement"`
  	Name        string			`json:"name" form:"name"`
  	Description string			`json:"description" form:"description"`
  	Icon        string			`json:"icon" form:"icon"`
  	Types       []string		`json:"types" form:"types" gorm:"type:text[]"`
  	CreatedAt   time.Time      	`json:"createdAt"`
  	UpdatedAt   time.Time      	`json:"updatedAt"`
  	GameProduct []GameProduct	`gorm:"foreignKey:gameId"`
  	Transaction []Transaction	`gorm:"foreignKey:gameId"`
}