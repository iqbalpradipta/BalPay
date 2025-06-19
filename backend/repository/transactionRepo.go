package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	Create(data *model.Transaction) error
	FindAll() ([]model.Transaction, error)
	FindById(id uint) (model.Transaction, error)
	Update(id uint, update *model.Transaction) error
	Delete(id uint) error
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{db}
}

func (t *transactionRepo) Create(data *model.Transaction) error {
	return t.db.Create(&data).Error

}

func (t *transactionRepo) FindAll() ([]model.Transaction, error) {
	var data []model.Transaction

	err := t.db.Preload("User").Preload("ProductDetail").Find(&data).Error

	return data, err
}

func (t *transactionRepo) FindById(id uint) (model.Transaction, error) {
	var data model.Transaction

	err := t.db.Preload("User").Preload("ProductDetail").First(&data, id).Error

	return data, err
}

func (t *transactionRepo) Update(id uint, update *model.Transaction) error {
	var data model.Transaction

	err := t.db.First(&data, id).Error; if err != nil {
		return err
	}

	if update.TransactionCode != "" {
		data.TransactionCode = update.TransactionCode
	}

	if update.GameUserId != "" {
		data.GameUserId = update.GameUserId
	}

	if update.StatusTransaction != "" {
		data.StatusTransaction = update.StatusTransaction
	}

	return t.db.Save(&update).Error
}

func (t *transactionRepo) Delete(id uint) error {
	return t.db.Delete(&model.Transaction{}, id).Error
}




