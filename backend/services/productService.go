package services

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetAllProduct() ([]model.Product, error)
	GetByIdProduct(id uint) (model.Product, error)
	UpdateProduct(id uint, update *model.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepo
}

func NewProductRepository(r repository.ProductRepo) ProductService {
	return &productService{r}
}

func (p *productService) CreateProduct(product *model.Product) error {
	return p.repo.Create(product)
}

func (p *productService) GetAllProduct() ([]model.Product, error) {
	return p.repo.FindAll()
}

func (p *productService) GetByIdProduct(id uint) (model.Product, error) {
	return p.repo.FindById(id)
}

func (p *productService) UpdateProduct(id uint, update *model.Product) error {
	return p.repo.Update(id, update)
}

func (p *productService) DeleteProduct(id uint) error {
	return p.repo.Delete(id)
}