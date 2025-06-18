package routes

import (
	"github.com/iqbalpradipta/BalPay/tree/main/backend/controllers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/middlewares"
	"github.com/labstack/echo/v4"
)

func ProductDetailRoutes(e *echo.Group, pc *controllers.ProductDetailController) {
	e.GET("/product_detail", pc.GetAllProductDetail, middlewares.Authentication("admin", "user"))
	e.GET("/product_detail/:id", pc.GetByIdProductDetail, middlewares.Authentication("admin", "user"))
	e.POST("/product_detail", pc.CreateProductDetail, middlewares.Authentication("admin", "user"))
	e.PUT("/product_detail/:id", pc.UpdateProductDetail, middlewares.Authentication("admin", "user"))
	e.DELETE("/product_detail/:id", pc.DeleteProductDetail, middlewares.Authentication("admin", "user"))
}