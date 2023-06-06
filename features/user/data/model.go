package data

import (
	"github.com/iqbalpradipta/BalPay/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name		string
	Username	string
	Email		string `gorm:"unique"`
	Password	string
	Role		string
}

func fromCore(dataCore user.Core) User {
	return User{
		Name: dataCore.Name,
		Username: dataCore.Name,
		Email: dataCore.Email,
		Password: dataCore.Password,
		Role : dataCore.Role,
	}
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID: data.ID,
		Name: data.Name,
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
		Role: data.Role,
	}
}

func toCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}