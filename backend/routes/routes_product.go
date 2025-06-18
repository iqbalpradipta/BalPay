package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/labstack/echo/v4"
)



func ProductRoutes(e *echo.Group, pc *controllers.ProductController) {
	e.GET("/product", pc.GetAllProduct, middlewares.Authentication("admin", "user"))
	e.GET("/product/:id", pc.GetByIdProduct, middlewares.Authentication("admin", "user"))
	e.POST("/product", pc.CreateProduct, middlewares.Authentication("admin", "user"))
	e.PUT("/product/:id", pc.UpdateProduct, middlewares.Authentication("admin", "user"))
	e.DELETE("/product/:id", pc.DeleteProduct, middlewares.Authentication("admin", "user"))
}