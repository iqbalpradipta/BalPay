package middlewares

import (
	"net/http"

	"github.com/iqbalpradipta/BalPay/utils/helper"
	"github.com/labstack/echo/v4"
)

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := ExtractToken(c)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("you are not an admin"))
		}
		return next(c)
	}
}

func isMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := ExtractToken(c)
		if role != "member" {
			return c.JSON(http.StatusForbidden, helper.FailedResponseHelper("please register to become a member"))
		}
		return next(c)
	}
}

