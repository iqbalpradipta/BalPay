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

type transactionPaymentService struct {
	transactionPaymentRepository service.TransactionPayment
}

func TransactionPaymentControllers(transactionPaymentRepository service.TransactionPayment ) *transactionPaymentService {
	return &transactionPaymentService{transactionPaymentRepository}
}

func(tp *transactionPaymentService) GetTransactionPayment(c echo.Context) error {
	transactionPayment, err := tp.transactionPaymentRepository.GetTransactionPayment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to get data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success to get data", transactionPayment))
}

func(tp *transactionPaymentService) GetTransactionPaymentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionPayment, err := tp.transactionPaymentRepository.GetTransactionPaymentById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to get data by id"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success to get data with id", transactionPayment))
}

func(tp *transactionPaymentService) CreateTransactionPayment(c echo.Context) error {
	transactionPayment := new(model.TransactionPayment)
	err := c.Bind(&transactionPayment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to bind data"))
	}

	transactionPaymentModel := model.TransactionPayment{
		PaymentStatus: transactionPayment.PaymentStatus,
		PaidAt: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		TransactionId: transactionPayment.TransactionId,
		PaymentMethodId: transactionPayment.PaymentMethodId,
	}

	response, err := tp.transactionPaymentRepository.CreateTransactionPayment(transactionPaymentModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to create data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success to create data", response))
}

func(tp *transactionPaymentService) UpdateTransactionPayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionPaymentData, err := tp.transactionPaymentRepository.GetTransactionPaymentById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to find id"))
	}

	transactionPayment := new(model.TransactionPayment)
	err = c.Bind(transactionPayment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to Bind data"))
	}

	if transactionPayment.PaymentStatus != "" {
		transactionPaymentData.PaymentStatus = transactionPayment.PaymentStatus
	}
	
	transactionPayment.UpdatedAt = time.Now()
	response, err := tp.transactionPaymentRepository.UpdateTransactionPayment(transactionPaymentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to update data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success to update data", response))
}

func(tp *transactionPaymentService) DeleteTransactionPayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionPaymentData, err := tp.transactionPaymentRepository.GetTransactionPaymentById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to find id"))
	}

	response, err := tp.transactionPaymentRepository.DeleteTransactionPayment(transactionPaymentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to delete data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success to delete data", response))
}