package delivery

import (
	"net/http"

	"github.com/iqbalpradipta/BalPay/features/user"
	"github.com/iqbalpradipta/BalPay/middlewares"
	"github.com/iqbalpradipta/BalPay/utils/helper"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	UserUsecase user.IusecaseInterface
}

func NewHandler(e *echo.Echo, useCase user.IusecaseInterface) {
	handler := &UserDelivery{
		UserUsecase: useCase,
	}

	e.GET("/user", handler.GetUser, middlewares.JWTMiddleware(), middlewares.IsAdmin)
	
}

func (d *UserDelivery) GetUser(c echo.Context) error {
	result, err := d.UserUsecase.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCoreList(result)))
}