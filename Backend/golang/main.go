package main

import (
	"net/http"

	"github.com/iqbalpradipta/BalPay/golang/src/config"
	"github.com/iqbalpradipta/BalPay/golang/src/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	config.DBConfig()
	routes.Routes(e.Group("/api/v1"))


	e.Logger.Fatal(e.Start(":8080"))
}
