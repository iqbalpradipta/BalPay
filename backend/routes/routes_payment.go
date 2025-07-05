package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
)

func PaymentRoutes(e *echo.Group, pm *controllers.PaymentController) {
	e.GET("/transaction/:id/pay", pm.Pay)
	e.POST("/webhook/xendit", pm.XenditCallback)
}