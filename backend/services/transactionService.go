package services

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
)

type TransactionService interface {
	CreateTransaction(data *model.Transaction) error
	GetAllTransaction(userId uint) ([]model.Transaction, error)
	GetByIdTransaction(id uint) (model.Transaction, error)
	UpdateTransaction(id uint, update *model.Transaction) error
	DeleteTransaction(id uint) error
	GetByJSONTransaction(id uint) (model.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepo
}

func NewTransactionService(t repository.TransactionRepo) TransactionService {
	return &transactionService{t}
}

func (t *transactionService) CreateTransaction(data *model.Transaction) error {
	return t.repo.Create(data)
}

func (t *transactionService) GetAllTransaction(userId uint) ([]model.Transaction, error) {
	return t.repo.FindAll(userId)
}

func (t *transactionService) GetByIdTransaction(id uint) (model.Transaction, error) {
	return t.repo.FindById(id)
}

func (t *transactionService) GetByJSONTransaction(id uint) (model.Transaction, error) {
	return t.repo.FindById(id)
}

func (t *transactionService) UpdateTransaction(id uint, update *model.Transaction) error {
	return t.repo.Update(id, update)
}

func (t *transactionService) DeleteTransaction(id uint) error {
	return t.repo.Delete(id)
}



