package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
)

type PaymentController struct {
	services.PaymentService
}

func NewPaymentController(ps services.PaymentService) *PaymentController {
	return &PaymentController{ps}
}

func (pc *PaymentController) Pay(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Invalid Id Invoice", err)
	}

	url, err := pc.PaymentService.CreateInvoice(uint(id))
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create Invoice", err)
	}

	return c.Redirect(http.StatusFound, url)
}

func (pc *PaymentController) XenditCallback(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Invalid Payload", err)
	}

	err := pc.PaymentService.HandleXenditCallback(payload)
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to process payment callback", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Webhook Process Success", payload)
}
