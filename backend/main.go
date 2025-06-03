package main

import (
	"log"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/config"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/migration"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	routes.Routes(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}