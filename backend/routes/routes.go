package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	api := e.Group("/api/v1")

	// Depedency Product
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductRepository(productRepo)
	productController := controllers.NewProductController(productService)
	ProductRoutes(api, productController)
}