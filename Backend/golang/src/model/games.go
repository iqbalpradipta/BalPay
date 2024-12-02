package model

import "time"

type Games struct {
	Id          int          	`json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
  	Name        string			`json:"name,omitempty" form:"name" gorm:"column:name"`
  	Description string			`json:"description,omitempty" form:"description" gorm:"column:description"`
  	Icon        string			`json:"icon,omitempty" form:"icon" gorm:"column:icon"`
  	Types       string			`json:"types,omitempty" form:"types" gorm:"column:types"`
  	CreatedAt   time.Time     	`json:"createdAt,omitempty" gorm:"column:createdAt"`
    UpdatedAt   time.Time     	`json:"updatedAt,omitempty" gorm:"column:updatedAt"`
  	GameProduct []GameProduct	`gorm:"foreignKey:gameId" json:"GameProduct,omitempty"`
  	Transaction []Transaction	`gorm:"foreignKey:gameId" json:"Transaction,,omitempty"`
}

func(Games) TableName() string {
	return "Games"
}