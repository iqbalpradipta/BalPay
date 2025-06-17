package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group, c *controllers.TransactionControllers) {
	e.GET("/transaction", c.GetAllTransaction)
	e.GET("/transaction/:id", c.GetByIdTransaction)
	e.POST("/transaction", c.CreateTransaction)
	e.PUT("/transaction/:id", c.UpdateTransaction)
	e.DELETE("/transaction/:id", c.DeleteTransaction)
}