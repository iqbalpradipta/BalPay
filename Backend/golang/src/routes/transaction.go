package routes

import (
	"github.com/iqbalpradipta/BalPay/golang/src/config"
	"github.com/iqbalpradipta/BalPay/golang/src/controllers"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/iqbalpradipta/BalPay/golang/src/middlewares"
	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	serviceTransaction := service.TransactionRepository(config.DB)
	transactionPaymentRepo := service.PaymentMethodRepository(config.DB)
	midtransService := service.NewMidtransService(transactionPaymentRepo) 
	controllerTransaction := controllers.TransactionController(serviceTransaction, midtransService)

	e.GET("/transaction", controllerTransaction.GetTransaction)
	e.GET("/transaction/:id", controllerTransaction.GetTransactionById)
	e.POST("/transaction", controllerTransaction.CreateTransaction, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.PUT("/transaction/:id", controllerTransaction.UpdateTransaction, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.DELETE("/transaction/:id", controllerTransaction.DeleteTransaction, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
}