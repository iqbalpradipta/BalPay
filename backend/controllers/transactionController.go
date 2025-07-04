package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/helpers"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/services"
	"github.com/labstack/echo/v4"
)

type TransactionControllers struct {
	services.TransactionService
}

func NewTransactionController(t services.TransactionService) *TransactionControllers  {
	return &TransactionControllers{t}
}

func (t *TransactionControllers) CreateTransaction(c echo.Context) error  {
	var data model.Transaction

	randomCode := uuid.New()

	if err := c.Bind(&data); err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	userID, ok := c.Get("user_id").(uint)
	log.Println(userID)
	if !ok {
		return c.JSON(http.StatusInternalServerError, echo.Map{"Messages": "Invalid UserID"})
	}
	data.TransactionCode = randomCode.String()
	data.UserID = userID
	data.StatusTransaction = "pending"


	err := t.TransactionService.CreateTransaction(&data)
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create data", err)
	}

	result, err := t.TransactionService.GetByIdTransaction(data.ID)
	if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to create data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Create Transaction", result)

}

func (t *TransactionControllers) GetAllTransaction(c echo.Context) error {
	data, err := t.TransactionService.GetAllTransaction(); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get All Transaction", data)
}

func (t *TransactionControllers) GetByIdTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := t.TransactionService.GetByIdTransaction(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to get data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Get Transaction By id", data)
}

func (t *TransactionControllers) UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := t.TransactionService.GetByIdTransaction(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to find data", err)
	}

	err = c.Bind(&data); if err != nil {
		return helpers.FailedResponse(c, http.StatusBadRequest, "Failed to bind data", err)
	}

	err = t.TransactionService.UpdateTransaction(uint(id), &data); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to update data", err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, "Success Update Transaction", data)
}

func (t *TransactionControllers) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := t.TransactionService.DeleteTransaction(uint(id)); if err != nil {
		return helpers.FailedResponse(c, http.StatusInternalServerError, "Failed to delete data", err)
	} 

	return helpers.SuccessResponse(c, http.StatusOK, "Success Delete Transaction", model.Transaction{})
}