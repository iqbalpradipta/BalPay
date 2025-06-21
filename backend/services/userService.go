package services

import (
	"strconv"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(data *model.User) error
	LoginUser(data *model.User) (string, error)
	GetAllUser() ([]model.User, error)
	GetByIdUser(id uint) (model.User, error)
	GetEmailUser(email string) (model.User, error)
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return err
	}

	data.Role = "User"
	data.Password = string(hashPassword)
	
	return u.userRepo.Create(data)
}


func (u *userService) LoginUser(data *model.User) (string, error) {
	user, err := u.userRepo.FindByEmail(data.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", err
	}

	id := strconv.Itoa(int(user.ID))

	token, err := middlewares.CreateToken(id, user.Name, user.Email, user.Role); if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userService) GetEmailUser(email string) (model.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userService) GetAllUser() ([]model.User, error) {
	return u.userRepo.FindAll()
}

func (u *userService) GetByIdUser(id uint) (model.User, error) {
	return u.userRepo.FindById(id)
}

func (u *userService) UpdateUser(id uint, update *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(update.Password), 14)
	if err != nil {
		return err
	}

	update.Password = string(hashPassword)
	return u.userRepo.Update(id, update)
}

func (u *userService) DeleteUser(email string) error {
	return u.userRepo.Delete(email)
}
