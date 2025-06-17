package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group, c *controllers.UserController) {
	e.GET("/user", c.GetAllUser)
	e.GET("/user/:id", c.GetByIdUser)
	e.POST("/user", c.CreateUser)
	e.PUT("/user/:id", c.UpdateUser)
	e.DELETE("/user/:email", c.DeleteUser)
}