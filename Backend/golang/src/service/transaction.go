package service

import (
    "github.com/iqbalpradipta/BalPay/golang/src/model"
    "gorm.io/gorm"
)

type Transaction interface {
    GetTransaction() ([]model.Transaction, error)
    GetTransactionById(id int) (model.Transaction, error)
    CreateTransaction(transaction model.Transaction) (model.Transaction, error)
    UpdateTransaction(transaction model.Transaction) (model.Transaction, error)
    DeleteTransaction(transaction model.Transaction) (model.Transaction, error)
}

func TransactionRepository(db *gorm.DB) *Repository {
    return &Repository{db}
}

func (r *Repository) GetTransaction() ([]model.Transaction, error) {
    var transactions []model.Transaction
    err := r.db.Preload("Users").Preload("Games").Preload("GameProduct").Preload("TransactionPayment").Find(&transactions).Error
    return transactions, err
}

func (r *Repository) GetTransactionById(id int) (model.Transaction, error) {
    var transaction model.Transaction
    err := r.db.Preload("Users").Preload("Games").Preload("GameProduct").Preload("TransactionPayment").First(&transaction, id).Error
    return transaction, err
}

func (r *Repository) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
    err := r.db.Create(&transaction).Error
    if err != nil {
        return transaction, err
    }
    err = r.db.Preload("Users").Preload("Games").Preload("GameProduct").Preload("TransactionPayment").First(&transaction, transaction.Id).Error
    return transaction, err
}

func (r *Repository) UpdateTransaction(transaction model.Transaction) (model.Transaction, error) {
    err := r.db.Save(&transaction).Error
    return transaction, err
}

func (r *Repository) DeleteTransaction(transaction model.Transaction) (model.Transaction, error) {
    err := r.db.Delete(&transaction).Error
    return transaction, err
}