package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	Create(payment *model.Payment) error
	FindByToken(token string) (*model.Payment, error)
	Update(payment *model.Payment) error
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db}
}

func (r *paymentRepo) Create(payment *model.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepo) FindByToken(token string) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Where("payment_token = ?", token).First(&payment).Error
	return &payment, err
}

func (r *paymentRepo) Update(payment *model.Payment) error {
	return r.db.Save(payment).Error
}