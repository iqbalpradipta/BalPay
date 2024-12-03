package service

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/iqbalpradipta/BalPay/golang/src/model"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransService interface {
	CreateMidtransTransaction(transaction model.Transaction) (*snap.Response, error)
	CreatePaymentWithToken(transaction model.Transaction) (model.TransactionPayment, *snap.Response, error)
}

type midtransService struct {
	paymentRepo TransactionPayment
}

func NewMidtransService(paymentRepo TransactionPayment) MidtransService {
	return &midtransService{paymentRepo: paymentRepo}
}

func (ms *midtransService) CreateMidtransTransaction(transaction model.Transaction) (*snap.Response, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return nil, errors.New("midtrans server key is not set")
	}

	snapClient := snap.Client{}
	snapClient.New(serverKey, midtrans.Sandbox)

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