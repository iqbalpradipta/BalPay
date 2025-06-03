package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string `json:"name" form:"name" gorm:"not null"`
	Email		string `json:"email" form:"email" gorm:"unique;not null"`
	Password	string `json:"password" form:"password" gorm:"not null"`
	Role		string `json:"role" form:"role" gorm:"default:user"`

	Transaction	[]Transaction	`gorm:"foreignKey:UserID"`
}