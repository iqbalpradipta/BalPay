package delivery

import (
	"net/http"
	"strconv"

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

	e.GET("/user", handler.GetUser, middlewares.JWTMiddleware())
	e.GET("/user/:id", handler.GetUserById, middlewares.JWTMiddleware())
	e.PUT("/user", handler.UpdateData, middlewares.JWTMiddleware(), middlewares.IsMember, middlewares.IsAdmin)
	e.POST("/user", handler.CreateData, middlewares.JWTMiddleware())
	e.DELETE("user/:id", handler.CreateData, middlewares.JWTMiddleware(), middlewares.IsAdmin)
}

func (d *UserDelivery) GetUser(c echo.Context) error {
	result, err := d.UserUsecase.GetUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("success get data", fromCoreList(result)))
}


func (d *UserDelivery) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error param"))
	}
	result, err := d.UserUsecase.GetUserById(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseDataHelper("get data success", fromCoreToResponse(result)))
}

func (d *UserDelivery) CreateData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}
	row, err := d.UserUsecase.CreateData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success create data"))
}

func (d *UserDelivery) UpdateData(c echo.Context) error {
	idToken, _  := middlewares.ExtractToken(c)
	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error request"))
	}
	row, err := d.UserUsecase.UpdateData(idToken, toCore(dataUpdate))
	if err != nil || row == 0{
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("update data success"))
}

func (d *UserDelivery) DeleteData(c echo.Context) error {
	idToken, _ := middlewares.ExtractToken(c)
	var dataDelete UserRequest
	row, err := d.UserUsecase.DeleteData(idToken, toCore(dataDelete))
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to delete data"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("success delete data"))
	
}