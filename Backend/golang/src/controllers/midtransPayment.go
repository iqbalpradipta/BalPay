package controllers

import (
	"net/http"

	"github.com/iqbalpradipta/BalPay/golang/src/helpers"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/labstack/echo/v4"
)

type MidtransPaymentController struct {
	midtransService service.MidtransService
}

func NewMidtransPaymentController(midtransService service.MidtransService) *MidtransPaymentController {
	return &MidtransPaymentController{midtransService: midtransService}
}

func (c *MidtransPaymentController) HandleCallback(ctx echo.Context) error {
	var notification map[string]interface{}

	if err := ctx.Bind(&notification); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when bind data"))
	}

	transaction, err := c.midtransService.ProcessNotification(notification)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Callback process", transaction))
}