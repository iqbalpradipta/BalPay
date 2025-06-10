package controllers

import (
	"net/http"
	"strconv"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/labstack/echo/v4"
)

type PaymentMethodControllers struct {
	services.PaymentMethod
}

func NewPaymentMethodControllers(s services.PaymentMethod) *PaymentMethodControllers  {
	return &PaymentMethodControllers{s}
}

func (pm *PaymentMethodControllers) CreatePaymentMethod(c echo.Context) error {
	var paymentMethod model.PaymentMethod
	
	err := c.Bind(&paymentMethod); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	err = pm.PaymentMethod.CreatePaymentMethod(&paymentMethod); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create payment method", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success create payment method", paymentMethod)
}

func (pm *PaymentMethodControllers) GetAllPaymentMethod(c echo.Context) error {
	data, err := pm.PaymentMethod.GetAllPaymentMethod(); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get Payment method", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success get all Payment Method", data)
}

func (pm *PaymentMethodControllers) GetByIdPaymentMethod(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := pm.PaymentMethod.GetByIdPaymentMethod(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get payment method by id", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success get payment Method", data)

}

func (pm *PaymentMethodControllers) UpdatePaymentMethod(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := pm.PaymentMethod.GetByIdPaymentMethod(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get Payment method", err)
	}

	err = c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}


	err = pm.PaymentMethod.UpdatePaymentMethod(uint(id), &data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to update payment method", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success to update payment method", data)
}

func (pm *PaymentMethodControllers) DeletePaymentMethod(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := pm.PaymentMethod.DeletePaymentMethod(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed delete payment method", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success delete data", model.PaymentMethod{})
}