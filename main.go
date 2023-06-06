package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/iqbalpradipta/BalPay/config"
	"github.com/iqbalpradipta/BalPay/factory"
	"github.com/iqbalpradipta/BalPay/migration"
	"github.com/iqbalpradipta/BalPay/utils/database/postgresql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidation struct {
	validator *validator.Validate
}

func (cv *CustomValidation) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main()  {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	cfg := config.GetConfig()
	db := postgresql.InitDB(cfg)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = &CustomValidation{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderWWWAuthenticate, echo.HeaderAccessControlAllowHeaders},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	e.Use(middleware.Recover())
	migration.InitMigrate(db)
	factory.InitFactory(e, db)
	err = e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
