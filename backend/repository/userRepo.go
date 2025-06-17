package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *model.User) error
	FindAll() ([]model.User, error)
	FindById(id uint) (model.User, error)
	Update(id uint, update *model.User) error
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(data *model.User) error {
	return u.db.Create(&data).Error
}

func (u *userRepo) FindAll() ([]model.User, error) {
	var data []model.User
	err := u.db.Find(&data).Error;
	return data, err
}

func (u *userRepo) FindById(id uint) (model.User, error) {
	var data model.User
	err := u.db.First(id, &data).Error

	return data, err
}

func (u *userRepo) Update(id uint, update *model.User) error {
	var data model.User

	err := u.db.First(id, &data).Error; if err != nil {
		return err
	}

	if update.Name != "" {
		data.Name = update.Name
	}

	if update.Email != "" {
		data.Email = update.Email
	}

	if update.Password != "" {
		data.Password = update.Password
	}

	if update.Role != "" {
		data.Role = update.Role
	}

	return u.db.Save(&update).Error
}

func (u *userRepo) Delete(id uint) error {
	return u.db.Delete(id).Error
}