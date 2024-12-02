package routes

import (
	"github.com/iqbalpradipta/BalPay/golang/src/config"
	"github.com/iqbalpradipta/BalPay/golang/src/controllers"
	"github.com/iqbalpradipta/BalPay/golang/src/middlewares"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/labstack/echo/v4"
)

func TransactionPaymentRoute(e *echo.Group) {
	serviceTransactionPayment := service.TransactionPaymentRepository(config.DB)
	controllerTransactionPayment := controllers.TransactionPaymentControllers(serviceTransactionPayment)

	e.GET("/transactionPayment", controllerTransactionPayment.GetTransactionPayment)
	e.GET("/transactionPayment/:id", controllerTransactionPayment.GetTransactionPaymentById)
	e.POST("/transactionPayment", controllerTransactionPayment.CreateTransactionPayment, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.PUT("/transactionPayment/:id", controllerTransactionPayment.UpdateTransactionPayment, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.DELETE("/transactionPayment/:id", controllerTransactionPayment.DeleteTransactionPayment, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
}