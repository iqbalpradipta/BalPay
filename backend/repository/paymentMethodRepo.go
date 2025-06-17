package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type PaymentMethod interface {
	Create(data *model.PaymentMethod) error
	FindAll() ([]model.PaymentMethod, error)
	FindById(id uint) (model.PaymentMethod, error)
	Update(id uint, update *model.PaymentMethod) error
	Delete(id uint) error
}

type paymentMethod struct {
	db *gorm.DB
}

func NewPaymentMethodRepo(db *gorm.DB) PaymentMethod {
	return paymentMethod{db}
}

func (p paymentMethod) Create(data *model.PaymentMethod) error {
	return p.db.Create(&data).Error;
}

func (p paymentMethod) FindAll() ([]model.PaymentMethod, error) {
	var data []model.PaymentMethod

	err := p.db.Find(&data).Error

	return data, err
}

func (p paymentMethod) FindById(id uint) (model.PaymentMethod, error) {
	var data model.PaymentMethod

	err := p.db.First(&data, id).Error;

	return data, err
}

func (p paymentMethod) Update(id uint, update *model.PaymentMethod) error {
	var data model.PaymentMethod
	if err := p.db.First(&data, id).Error; err != nil {
		return err
	}

	if update.Name != "" {
		data.Name = update.Name
	}

	return p.db.Save(&update).Error
}

func (p paymentMethod) Delete(id uint) error {
	return p.db.Delete(model.PaymentMethod{}, &id).Error
}