package services

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
)

type ProductDetailService interface {
	CreateProductDetail(data *model.ProductDetail) error
	GetAllProductDetail() ([]model.ProductDetail, error)
	GetByIdProductDetail(id uint) (model.ProductDetail, error)
	UpdateProductDetail(id uint, data *model.ProductDetail) error
	DeleteProductDetail(id uint) error
}

type productDetailService struct {
	repo repository.ProductDetailRepo
}

func NewProductDetailService(r repository.ProductDetailRepo) ProductDetailService {
	return &productDetailService{r}
}

func (p *productDetailService) CreateProductDetail(data *model.ProductDetail) error {
	return p.repo.Create(data)
}

func (p *productDetailService) DeleteProductDetail(id uint) error {
	return p.repo.Delete(id)
}

func (p *productDetailService) GetAllProductDetail() ([]model.ProductDetail, error) {
	return p.repo.FindAll()
}

func (p *productDetailService) GetByIdProductDetail(id uint) (model.ProductDetail, error) {
	return p.repo.FindById(id)
}

func (p *productDetailService) UpdateProductDetail(id uint, data *model.ProductDetail) error {
	return p.repo.Update(id, data)
}

