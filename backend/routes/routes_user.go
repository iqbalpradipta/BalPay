package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group, c *controllers.UserController) {
	e.GET("/user", c.GetAllUser, middlewares.Authentication("admin", "user"))
	e.GET("/user/:id", c.GetByIdUser, middlewares.Authentication("admin", "user"))
	e.GET("/user/email", c.GetEmailUser, middlewares.Authentication("admin", "user"))
	e.POST("/login", c.LoginUser)
	e.POST("/user", c.CreateUser)
	e.PUT("/user/:id", c.UpdateUser, middlewares.Authentication("admin", "user"))
	e.DELETE("/user/:email", c.DeleteUser, middlewares.Authentication("admin", "user"))
}