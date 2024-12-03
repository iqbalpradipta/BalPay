package service

import (
	"github.com/iqbalpradipta/BalPay/golang/src/model"
	"gorm.io/gorm"
)

type TransactionPayment interface {
	GetTransactionPayment() ([]model.TransactionPayment, error)
	GetTransactionPaymentById(id int) (model.TransactionPayment, error)
	CreateTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)
	UpdateTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)
	DeleteTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)
}

func TransactionPaymentRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func(r *Repository) GetTransactionPayment() ([]model.TransactionPayment, error) {
	var transactionPayment []model.TransactionPayment
	err := r.db.Preload("Transaction").Preload("PaymentMethod").Find(&transactionPayment).Error

	return transactionPayment, err
}

func(r *Repository) GetTransactionPaymentById(id int) (model.TransactionPayment, error)  {
	var transactionPayment model.TransactionPayment
	err := r.db.Preload("Transaction").Preload("PaymentMethod").First(&transactionPayment, id).Error

	return transactionPayment, err
}

func(r *Repository) CreateTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)  {
	err := r.db.Create(&transactionPayment).Error

	return transactionPayment, err
}

func(r *Repository) UpdateTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)  {
	err := r.db.Save(&transactionPayment).Error

	return transactionPayment, err
}

func(r *Repository) DeleteTransactionPayment(transactionPayment model.TransactionPayment) (model.TransactionPayment, error)  {
	err := r.db.Delete(&transactionPayment).Error

	return transactionPayment, err
}