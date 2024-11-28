package model

import "time"

type Users struct {
	Id            int    		`json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Name          string 		`json:"name" form:"name" gorm:"column:name"`
	Email         string 		`json:"email" form:"email" gorm:"column:email"`
	Password      string 		`json:"password" form:"password" gorm:"column:password"`
	PhoneNumber   string 		`json:"phoneNumber" form:"phoneNumber" gorm:"column:phoneNumber"`
	Role          string 		`gorm:"column:role;default:member"`
	PhotoProfile  string 		`json:"photoProfile" form:"photoProfile" gorm:"column:photoProfile"`
	CreatedAt     time.Time     `json:"createdAt" gorm:"column:createdAt"`
    UpdatedAt     time.Time     `json:"updatedAt" gorm:"column:updatedAt"`
	Transaction   []Transaction	`gorm:"foreignKey:userId"`	
}

func(Users) TableName() string {
	return "Users"
}