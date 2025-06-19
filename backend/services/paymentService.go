package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/iqbalpradipta/BalPay/tree/main/backend/model"
	"github.com/iqbalpradipta/BalPay/tree/main/backend/repository"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type PaymentService interface {
	CreateInvoice(transactionID uint) (string, error)
	HandleXenditCallback(payload map[string]interface{}) error 
}

type paymentService struct {
	transactionRepo repository.TransactionRepo
	paymentRepo     repository.PaymentRepo
}

func NewPaymentService(tr repository.TransactionRepo, pr repository.PaymentRepo) PaymentService {
	return &paymentService{tr, pr}
}

func (s *paymentService) CreateInvoice(transactionID uint) (string, error) {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_API_KEY")

	data, err := s.transactionRepo.FindById(transactionID)
	if err != nil {
		return "", err
	}

	params := invoice.CreateParams{
		ExternalID:  data.TransactionCode,
		Amount: float64(data.ProductDetail.Price),
		PayerEmail:  data.User.Email,
		Description: fmt.Sprintf("Pembelian %s oleh %s", data.ProductDetail.Name, data.User.Email),

	}

	inv, _ := invoice.Create(&params);
	expiredAt := inv.ExpiryDate

	payment := &model.Payment{
		TransactionID: data.ID,
		PaymentType:   "xendit_invoice",
		PaymentToken:  inv.ID,
		PaymentURL:    inv.InvoiceURL,
		Status:        inv.Status,
		ExpiredAt:     expiredAt,
	}

	if err := s.paymentRepo.Create(payment); err != nil {
		return "", err
	}
	log.Printf("Invoice created: %s\n", inv.InvoiceURL)

	return inv.InvoiceURL, nil
}

func (s *paymentService) HandleXenditCallback(payload map[string]interface{}) error {
	token, ok := payload["id"].(string)
	if !ok {
		return fmt.Errorf("invalid payload: token not found")
	}
	statusRaw, ok := payload["status"].(string)
	if !ok {
		return fmt.Errorf("invalid payload: status not found")
	}

	status := strings.ToLower(statusRaw)

	payment, err := s.paymentRepo.FindByToken(token)
	if err != nil {
		return err
	}

	payment.Status = status
	if status == "paid" {
		now := time.Now()
		payment.PaidAt = &now
	}

	if err := s.paymentRepo.Update(payment); err != nil {
		return err
	}
	transaction, err := s.transactionRepo.FindById(payment.TransactionID)
	if err != nil {
		return err
	}
	transaction.StatusTransaction = payment.Status

	return s.transactionRepo.Update(transaction.ID, &transaction)
}
