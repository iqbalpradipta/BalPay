package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(product *model.Product) error
	FindAll() ([]model.Product, error)
	FindById(id uint) (model.Product, error)
	Update(id uint, product *model.Product) error
	Delete(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepo {
	return &productRepo{db}
}

func(r *productRepo) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func(r *productRepo) FindAll() ([]model.Product, error) {
	var product []model.Product
	err := r.db.Find(&product).Error
	return product, err
}

func(r *productRepo) FindById(id uint) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepo) Update(id uint, update *model.Product) error {
	var data model.Product
	if err := r.db.First(&data, id).Error; err != nil {
		return err
	}

	if update.Name != "" {
		data.Name = update.Name
	}
	if update.Image != "" {
		data.Image = update.Image
	}
	if update.Description != "" {
		data.Description = update.Description
	}

	return r.db.Save(&data).Error
}

func (r *productRepo) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

