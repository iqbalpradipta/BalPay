package migration

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Transaction{},
		&model.ProductDetail{},
		&model.Product{},
		&model.PaymentMethod{},
		&model.Payment{},
	)
}