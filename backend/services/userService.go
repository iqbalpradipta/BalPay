package services

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(data *model.User) error
	GetAllUser() ([]model.User, error)
	GetByIdUser(id uint) (model.User, error)
	UpdateUser(id uint, update *model.User) error
	DeleteUser(email string) error
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(u repository.UserRepo) UserService {
	return &userService{u}
}

func (u *userService) CreateUser(data *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14);
	if err != nil {
		return err
	}

	data.Password = string(hashPassword)

	return u.userRepo.Create(data)
}

func (u *userService) GetAllUser() ([]model.User, error) {
	return u.userRepo.FindAll()
}

func (u *userService) GetByIdUser(id uint) (model.User, error) {
	return u.userRepo.FindById(id)
}

func (u *userService) UpdateUser(id uint, update *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(update.Password), 14);
	if err != nil {
		return err
	}

	update.Password = string(hashPassword)
	return u.userRepo.Update(id, update)
}

func (u *userService) DeleteUser(email string) error {
	return u.userRepo.Delete(email)
}



