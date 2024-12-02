package model

import "time"

type Users struct {
	Id            int    		`json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id"`
	Name          string 		`json:"name,omitempty" form:"name" gorm:"column:name"`
	Email         string 		`json:"email,omitempty" form:"email" gorm:"column:email"`
	Password      string 		`json:"password,omitempty" form:"password" gorm:"column:password"`
	PhoneNumber   string 		`json:"phoneNumber,omitempty" form:"phoneNumber" gorm:"column:phoneNumber"`
	Role          string 		`gorm:"column:role;default:member" json:"role,omitempty"`
	PhotoProfile  string 		`json:"photoProfile,omitempty" form:"photoProfile" gorm:"column:photoProfile"`
	CreatedAt     time.Time     `json:"createdAt,omitempty" gorm:"column:createdAt"`
    UpdatedAt     time.Time     `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	Transaction   []Transaction	`gorm:"foreignKey:userId" json:"Transaction,omitempty"`	
}

func(Users) TableName() string {
	return "Users"
}