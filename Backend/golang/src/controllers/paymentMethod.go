package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iqbalpradipta/BalPay/golang/src/helpers"
	"github.com/iqbalpradipta/BalPay/golang/src/model"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/labstack/echo/v4"
)

type paymentMethodService struct {
	paymentMethodRepository service.PaymentMethod
}

func PaymentController(paymentMethodRepository service.PaymentMethod) *paymentMethodService  {
	return &paymentMethodService{paymentMethodRepository}
}

func(ps *paymentMethodService) GetPaymentMethod(c echo.Context) error {
	paymentMethod, err := ps.paymentMethodRepository.GetPaymentMethod()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to get data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Get data", paymentMethod))
}

func(ps *paymentMethodService) GetPaymentMethodId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paymentMethod, err := ps.paymentMethodRepository.GetPaymentMethodId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to get data by id"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data by id", paymentMethod))
}

func(ps *paymentMethodService) CreatePaymentMethod(c echo.Context) error {
	paymentMethod := new(model.PaymentMethod)
	err := c.Bind(&paymentMethod)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed when bind data"))
	}

	paymentMethodModel := model.PaymentMethod {
		Name: paymentMethod.Name,
		Description: paymentMethod.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := ps.paymentMethodRepository.CreatePaymentMethod(paymentMethodModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to created data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success create data", data))
}