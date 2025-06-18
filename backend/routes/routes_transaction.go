package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group, c *controllers.TransactionControllers) {
	e.GET("/transaction", c.GetAllTransaction, middlewares.Authentication("admin", "user"))
	e.GET("/transaction/:id", c.GetByIdTransaction, middlewares.Authentication("admin", "user"))
	e.POST("/transaction", c.CreateTransaction, middlewares.Authentication("admin", "user"))
	e.PUT("/transaction/:id", c.UpdateTransaction, middlewares.Authentication("admin", "user"))
	e.DELETE("/transaction/:id", c.DeleteTransaction, middlewares.Authentication("admin", "user"))
}