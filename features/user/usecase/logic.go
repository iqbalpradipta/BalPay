package usecase

import (
	"errors"

	"github.com/iqbalpradipta/BalPay/features/user"
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
	if data.ID == 0 {
		return data, errors.New("id not found")
	}
	result, err := u.userData.SelectDataById(id)
	return result, err
}

func (u *userUseCase) CreateData(data user.Core) (int, error) {
	if data.Name == ""{
		return -1, errors.New("name must be fill")
	}
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