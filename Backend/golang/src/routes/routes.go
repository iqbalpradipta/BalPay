package routes

import "github.com/labstack/echo/v4"

func Routes(e *echo.Group) {
	TransactionRoutes(e)
	PaymentMethodRoutes(e)
}