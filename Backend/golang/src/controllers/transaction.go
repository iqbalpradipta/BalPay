package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iqbalpradipta/BalPay/golang/src/helpers"
	"github.com/iqbalpradipta/BalPay/golang/src/middlewares"
	"github.com/iqbalpradipta/BalPay/golang/src/model"
	"github.com/iqbalpradipta/BalPay/golang/src/service"
	"github.com/labstack/echo/v4"
)

type transactionService struct {
	transactionRepository service.Transaction
}

func TransactionController(transactionRepository service.Transaction) *transactionService {
	return &transactionService{transactionRepository}
}

func (ts *transactionService) GetTransaction(c echo.Context) error {
	transaction, err := ts.transactionRepository.GetTransaction()
	if(err != nil) {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error when get data"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data transaction", transaction))
}

func (ts *transactionService) GetTransactionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := ts.transactionRepository.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed get data transaction by Id"))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success get data transaction by id", transaction))
}

func(ts *transactionService) CreateTransaction(c echo.Context) error {
	userId, _ := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("Invalid token"))
	} 
	
	transaction := new(model.Transaction)
	err := c.Bind(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed when bind data"))
	}

	transactionModel := model.Transaction{
		Status: transaction.Status,
		Amount: transaction.Amount,
		UserId: userId,
		GameId: transaction.GameId,
		GameProductId: transaction.GameProductId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdTransaction, err := ts.transactionRepository.CreateTransaction(transactionModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Error create data transaction"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success create data", createdTransaction))
}

func (ts *transactionService) UpdateTransaction(c echo.Context) (err error) {
	userId, _ := middlewares.ExtractToken(c)
	if userId == -1 {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("Invalid token"))
	} 

	id, _ := strconv.Atoi(c.Param("id"))
	transactionData, err := ts.transactionRepository.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("Failed to found id"))
	}

	transaction:= new(model.Transaction)
	err = c.Bind(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to bind data"))
	}

	if transaction.Id != 0 {
		transactionData.Id = transaction.Id
	}

	if transaction.Status != "" {
		transactionData.Status = transaction.Status
	}

	if transaction.Amount != 0 {
		transactionData.Amount = transaction.Amount
	}

	transactionData.UpdatedAt = time.Now()
	response, err := ts.transactionRepository.UpdateTransaction(transactionData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to update data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success Update Data", response))
}

func (ts *transactionService) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transactionData, err := ts.transactionRepository.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Failed to get data with id"))
	}

	data, err := ts.transactionRepository.DeleteTransaction(transactionData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Failed to delete data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success delete data", data))
}