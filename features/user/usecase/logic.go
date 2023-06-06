package usecase

import (
	"errors"

	"github.com/iqbalpradipta/BalPay/features/user"
	"github.com/iqbalpradipta/BalPay/middlewares"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData user.IdataInterface
}

func NewLogic(data user.IdataInterface) user.IusecaseInterface {
	return &userUseCase {
		userData: data,
	}
}

func (u *userUseCase) GetUser() ([]user.Core, error) {
	result, err := u.userData.SelectAllData()
	return  result, err
}

func (u *userUseCase) GetUserById(id int) (data user.Core, err error) {
	data, err = u.userData.SelectDataById(id)
	if err != nil {
		return data, err
	}
	return data, nil
}



func (u *userUseCase) CreateData(data user.Core) (int, error) {
	passBcrypt := []byte(data.Password)
	hash, errHash := bcrypt.GenerateFromPassword(passBcrypt, bcrypt.DefaultCost)
	if errHash != nil {
		return -2, errors.New("failed to hashing password")
	}
	if data.Name == ""{
		return -1, errors.New("name must be fill")
	}
	if data.Password == "" {
		return -1, errors.New("password must be fill")
	}
	data.Password = string(hash)
	row, err := u.userData.InsertData(data)
	return row, err
}

func (u *userUseCase) UpdateData(id int, updateData user.Core) (int, error) {
	if updateData.ID == 0 {
		return -1, errors.New("id not found")
	}
	row, err := u.userData.PostData(id, updateData)
	return row, err
}

func (u *userUseCase) DeleteData(id int, deleteData user.Core) (int, error) {
	if deleteData.ID == 0 {
		return -1, errors.New("id not found")
	}
	row, err := u.userData.DelData(id, deleteData)
	return row, err
}

func (u *userUseCase) LoginData(email, password string) (string, string, string) {
	if email == "" || password == "" {
		return "please input email and password", "", ""
	}
	row, errEmail := u.userData.AuthData(email)
	if errEmail != nil {
		return "email not found", "", ""
	}

	errAuth := bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(password))
	if errAuth != nil {
		return "wrong Password", "", ""
	}

	token, errToken := middlewares.CreateToken(int(row.ID), row.Role)

	if errToken != nil {
		return "error to Create Token", "", ""
	}
	return token, row.Role, row.Username
}