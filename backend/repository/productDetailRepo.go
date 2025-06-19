package repository

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"gorm.io/gorm"
)

type ProductDetailRepo interface {
	Create(data *model.ProductDetail) error
	FindAll() ([]model.ProductDetail, error)
	FindById(id uint) (model.ProductDetail, error)
	Update(id uint, update *model.ProductDetail) error
	Delete(id uint) error
}

type productDetailRepo struct {
	db *gorm.DB
}

func NewProductDetailRepo(db *gorm.DB) ProductDetailRepo {
	return &productDetailRepo{db}
}

func (r *productDetailRepo) Create(data *model.ProductDetail) error {
	return r.db.Preload("Product").Create(&data).Error
}

func (r *productDetailRepo) FindAll() ([]model.ProductDetail, error) {
	var productDetail []model.ProductDetail

	err := r.db.Preload("Product").Find(&productDetail).Error

	return productDetail, err
}

func (r *productDetailRepo) FindById(id uint) (model.ProductDetail, error) {
	var productDetail model.ProductDetail

	err := r.db.Preload("Product").First(&productDetail, id).Error

	return productDetail, err
}

func (r *productDetailRepo) Update(id uint, update *model.ProductDetail) error {
	var data model.ProductDetail
	err := r.db.First(&data, id).Error; if err != nil {
		return err
	}

	
	if update.Name != "" {
		data.Name = update.Name
	}

	if update.Stock != 0 {
		data.Stock = update.Stock
	}

	if update.Price != 0 {
		data.Price = update.Price
	}

	return r.db.Save(&update).Error
}

func (r *productDetailRepo) Delete(id uint) error {
	return r.db.Delete(&model.ProductDetail{}, id).Error
}


