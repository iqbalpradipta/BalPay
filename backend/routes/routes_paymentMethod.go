package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/labstack/echo/v4"
)

func PaymentMethodRoutes(e *echo.Group, pm *controllers.PaymentMethodControllers) {
	e.GET("/paymentMethod", pm.GetAllPaymentMethod, middlewares.Authentication("admin", "user"))
	e.GET("/paymentMethod/:id", pm.GetByIdPaymentMethod, middlewares.Authentication("admin", "user"))
	e.POST("/paymentMethod", pm.CreatePaymentMethod, middlewares.Authentication("admin", "user"))
	e.PUT("/paymentMethod/:id", pm.UpdatePaymentMethod, middlewares.Authentication("admin", "user"))
	e.DELETE("/paymentMethod/:id", pm.DeletePaymentMethod, middlewares.Authentication("admin", "user"))
}