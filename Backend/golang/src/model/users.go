package model

import "time"

type Users struct {
	Id            int    		`json:"id" gorm:"primaryKey:autoIncrement"`
	Name          string 		`json:"name" form:"name"`
	Email         string 		`json:"email" form:"email"`
	Password      string 		`json:"password" form:"password"`
	PhoneNumber   string 		`json:"phoneNumber" form:"phoneNumber"`
	Role          string 		`gorm:"default:member"`
	PhotoProfile  string 		`json:"-" form:"photoProfile"`
	CreatedAt     time.Time   	`json:"createdAt"`
	UpdatedAt     time.Time   	`json:"updatedAt"`
	Transaction   []Transaction	`gorm:"foreignKey:userId"`	
}