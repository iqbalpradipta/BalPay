package services

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
)

type PaymentMethod interface {
	CreatePaymentMethod(data *model.PaymentMethod) error
	GetAllPaymentMethod() ([]model.PaymentMethod, error)
	GetByIdPaymentMethod(id uint) (model.PaymentMethod, error)
	UpdatePaymentMethod(id uint, data *model.PaymentMethod) error
	DeletePaymentMethod(id uint) error
}

type paymentMethod struct {
	repo repository.PaymentMethod
}


func NewPaymentMethodService(r repository.PaymentMethod) PaymentMethod {
	return &paymentMethod{r}
}


func (p *paymentMethod) CreatePaymentMethod(data *model.PaymentMethod) error {
	return p.repo.Create(data)
}

func (p *paymentMethod) GetAllPaymentMethod() ([]model.PaymentMethod, error) {
	return p.repo.FindAll()
}

func (p *paymentMethod) GetByIdPaymentMethod(id uint) (model.PaymentMethod, error) {
	return p.repo.FindById(id)
}

func (p *paymentMethod) UpdatePaymentMethod(id uint, data *model.PaymentMethod) error {
	return p.repo.Update(id, data)
}

func (p *paymentMethod) DeletePaymentMethod(id uint) error {
	return p.repo.Delete(id)
}

