package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/labstack/echo/v4"
)



func ProductRoutes(e *echo.Group, pc *controllers.ProductController) {
	e.GET("/product", pc.GetAllProduct)
	e.GET("/product/:id", pc.GetByIdProduct)
	e.POST("/product", pc.CreateProduct)
	e.PUT("/product/:id", pc.UpdateProduct)
	e.DELETE("/product/:id", pc.DeleteProduct)
}