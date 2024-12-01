package routes

import (
	"github.com/iqbalpradipta/BalPay/golang/src/config"
	"github.com/iqbalpradipta/BalPay/golang/src/controllers"
	"github.com/iqbalpradipta/BalPay/golang/src/middlewares"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/labstack/echo/v4"
)

func PaymentMethodRoutes(e *echo.Group) {
	servicePaymentMethod := service.PaymentMethodRepository(config.DB)
	controllerPaymentMethod := controllers.PaymentController(servicePaymentMethod)

	e.GET("/paymentMethod", controllerPaymentMethod.GetPaymentMethod)
	e.GET("/paymentMethod/:id", controllerPaymentMethod.GetPaymentMethodId)
	e.POST("/paymentMethod", controllerPaymentMethod.CreatePaymentMethod, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.PUT("/paymentMethod/:id", controllerPaymentMethod.UpdatePaymentMethod, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
	e.DELETE("/paymentMethod/:id", controllerPaymentMethod.DeletePaymentMethod, middlewares.JWTMiddleware(), middlewares.Authentication([]string{"admin"}))
}
