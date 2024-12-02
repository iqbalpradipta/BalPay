package service

import (
	"github.com/iqbalpradipta/BalPay/golang/src/model"
	"gorm.io/gorm"
)

type PaymentMethod interface {
	GetPaymentMethod() ([]model.PaymentMethod, error)
	GetPaymentMethodId(id int) (model.PaymentMethod, error)
	CreatePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	UpdatePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	DeletePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
}

func PaymentMethodRepository(db *gorm.DB) *Repository {
    return &Repository{db}
}

func(r *Repository) GetPaymentMethod() ([]model.PaymentMethod, error) {
	var paymentMethod []model.PaymentMethod
	err := r.db.Joins("TransactionPayment").Find(&paymentMethod).Error
	
	return paymentMethod, err
}

func(r *Repository) GetPaymentMethodId(id int) (model.PaymentMethod, error) {
	var paymentMethod model.PaymentMethod
	err := r.db.Joins("TransactionPayment").First(&paymentMethod, id).Error

	return paymentMethod, err
}

func(r *Repository) CreatePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	err := r.db.Create(&paymentMethod).Error

	return paymentMethod, err
}

func(r *Repository) UpdatePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	err := r.db.Save(&paymentMethod).Error

	return paymentMethod, err
}

func(r *Repository) DeletePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
	err := r.db.Delete(&paymentMethod).Error

	return paymentMethod, err
}