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

	// Depedency
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductRepository(productRepo)
	productController := controllers.NewProductController(productService)
	ProductRoutes(api, productController)

	productDetailRepo := repository.NewProductDetailRepo(db)
	productDetailService := services.NewProductDetailService(productDetailRepo)
	productDetailController := controllers.NewProductDetailController(productDetailService)
	ProductDetailRoutes(api, productDetailController)

	paymentMethodRepo := repository.NewPaymentMethodRepo(db)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepo)
	paymentMethodController := controllers.NewPaymentMethodControllers(paymentMethodService)
	PaymentMethodRoutes(api, paymentMethodController)

	transactionRepo := repository.NewTransactionRepo(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionControllers := controllers.NewTransactionController(transactionService)
	TransactionRoutes(api, transactionControllers)

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userControllers := controllers.NewUserController(userService)
	UserRoutes(api, userControllers)
}