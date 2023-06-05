package data

import (
	"errors"

	"github.com/iqbalpradipta/BalPay/features/user"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.IdataInterface {
	return &userData {
		db: db,
	}
}

func (repo *userData) SelectAllData() ([]user.Core, error) {
	var data []User
	tx := repo.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (repo *userData) SelectDataById(id int) (user.Core, error) {
	var data User
	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	dataCore := fromCore(user.Core{})
	return dataCore.toCore(), nil
}

func (repo *userData) InsertData(data user.Core) (int, error){
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, nil
	}
	return int(tx.RowsAffected), nil
}

func (repo *userData) PostData(id int, updatedData user.Core) (int, error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(fromCore(updatedData))
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}
	return int(tx.RowsAffected), nil
}
func (repo *userData) DelData(id int, deletedData user.Core) (int, error) {
	tx := repo.db.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete data")
	}
	return int(tx.RowsAffected), nil
}
