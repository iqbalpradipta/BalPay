package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/labstack/echo/v4"
)

func PaymentMethodRoutes(e *echo.Group, pm *controllers.PaymentMethodControllers) {
	e.GET("/paymentMethod", pm.GetAllPaymentMethod)
	e.GET("/paymentMethod/:id", pm.GetByIdPaymentMethod)
	e.POST("/paymentMethod", pm.CreatePaymentMethod)
	e.PUT("/paymentMethod/:id", pm.UpdatePaymentMethod)
	e.DELETE("/paymentMethod/:id", pm.DeletePaymentMethod)
}