package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	"github.com/iqbalpradipta/BalPay/golang/src/model"
)

type MidtransService interface {
	CreateMidtransTransaction(transaction model.Transaction) (*snap.Response, error)
	CreatePaymentWithToken(transaction model.Transaction) (model.TransactionPayment, *snap.Response, error)
	ProcessNotification(notification map[string]interface{}) (model.Transaction, error)
}

type midtransService struct {
	paymentRepo     TransactionPayment
	transactionRepo Transaction
	coreAPI         coreapi.Client
}

func NewMidtransService(paymentRepo TransactionPayment, transactionRepo Transaction) MidtransService {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		panic("MIDTRANS_SERVER_KEY is not set")
	}

	core := coreapi.Client{}
	core.New(serverKey, midtrans.Sandbox)

	return &midtransService{
		paymentRepo:     paymentRepo,
		transactionRepo: transactionRepo,
		coreAPI:         core,
	}
}

func (ms *midtransService) CreateMidtransTransaction(transaction model.Transaction) (*snap.Response, error) {
	snapClient := snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapResp, err := snapClient.CreateTransaction(snapReq)
	if err != nil {
		return nil, err
	}

	return snapResp, nil
}

func (ms *midtransService) CreatePaymentWithToken(transaction model.Transaction) (model.TransactionPayment, *snap.Response, error) {
	snapResp, err := ms.CreateMidtransTransaction(transaction)
	if err != nil {
		return model.TransactionPayment{}, nil, err
	}

	payment := model.TransactionPayment{
		PaymentStatus:   "pending",
		PaidAt:          time.Time{},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		TransactionId:   transaction.Id,
		PaymentMethodId: 2,
	}

	createdPayment, err := ms.paymentRepo.CreateTransactionPayment(payment)
	if err != nil {
		return model.TransactionPayment{}, nil, err
	}

	return createdPayment, snapResp, nil
}

func (ms *midtransService) ProcessNotification(notification map[string]interface{}) (model.Transaction, error) {
	orderIDStr, ok := notification["order_id"].(string)
	if !ok {
		return model.Transaction{}, errors.New("invalid order_id in notification")
	}

	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		return model.Transaction{}, errors.New("order_id is not a valid integer")
	}

	transaction, err := ms.transactionRepo.GetTransactionById(orderID)
	if err != nil {
		return model.Transaction{}, errors.New("transaction not found")
	}

	transactionStatus := notification["transaction_status"].(string)
	fraudStatus := notification["fraud_status"].(string)

	switch transactionStatus {
	case "capture":
		if fraudStatus == "accept" {
			transaction.Status = "success"
		} else if fraudStatus == "challenge" {
			transaction.Status = "challenge"
		}
	case "settlement":
		transaction.Status = "success"
	case "pending":
		transaction.Status = "pending"
	case "deny", "cancel", "expire":
		transaction.Status = "failed"
	default:
		transaction.Status = "unknown"
	}

	_, err = ms.transactionRepo.UpdateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}