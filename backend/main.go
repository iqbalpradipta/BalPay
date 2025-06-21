package main

import (
	"log"
	"net/http"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/config"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/migration"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	

	db := config.InitDB()

	if err := migration.AutoMigrate(db); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	routes.Routes(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}