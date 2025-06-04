package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/labstack/echo/v4"
)

func ProductDetailRoutes(e *echo.Group, pc *controllers.ProductDetailController) {
	e.GET("/product_detail", pc.GetAllProductDetail)
	e.GET("/product_detail/:id", pc.GetByIdProductDetail)
	e.POST("/product_detail", pc.CreateProductDetail)
	e.PUT("/product_detail/:id", pc.UpdateProductDetail)
	e.DELETE("/product_detail/:id", pc.DeleteProductDetail)
}