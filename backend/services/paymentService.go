package services

import (
	"fmt"
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

	inv, _ := invoice.Create(&params)
	expiredAt := inv.ExpiryDate

	payment := &model.Payment{
		TransactionID: data.ID,
		PaymentType:   "xendit_invoice",
		PaymentToken:  inv.ID,
		PaymentURL:    inv.InvoiceURL,
		Status:        inv.Status,
		ExpiredAt:     expiredAt,
		ExternalID:    data.TransactionCode,
	}


	if err := s.paymentRepo.Create(payment); err != nil {
		return "", err
	}

	return inv.InvoiceURL, nil
}

func (s *paymentService) HandleXenditCallback(payload map[string]interface{}) error {
    token, ok := payload["id"].(string)
    if !ok {
        return fmt.Errorf("invalid payload: id not found")
    }

    statusRaw, ok := payload["status"].(string)
    if !ok {
        return fmt.Errorf("invalid payload: status not found")
    }
    externalID, ok := payload["external_id"].(string)
    if !ok {
        return fmt.Errorf("invalid payload: external_id not found")
    }

    status := strings.ToLower(statusRaw)

    payment, err := s.paymentRepo.FindByToken(token)
    if err != nil {
        return fmt.Errorf("payment not found: %w", err)
    }

    if payment.ExternalID != externalID {
        return fmt.Errorf("mismatched external_id: payload %s, payment %s", externalID, payment.ExternalID)
    }

    payment.Status = status
    if status == "paid" {
        now := time.Now()
        payment.PaidAt = &now
    }

    if err := s.paymentRepo.Update(payment); err != nil {
        return fmt.Errorf("failed to update payment: %w", err)
    }

    transaction, err := s.transactionRepo.FindByCode(externalID)
    if err != nil {
        return fmt.Errorf("transaction not found: %w", err)
    }

    transaction.StatusTransaction = status

    return s.transactionRepo.Update(transaction.ID, transaction)
}

