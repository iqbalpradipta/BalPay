package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	TransactionID uint       `json:"transactionId"`                    
	PaymentType   string     `json:"paymentType"`                
	PaymentToken  string     `json:"paymentToken"`                    
	PaymentURL    string     `json:"paymentUrl"`                  
	Status        string     `json:"status"`  
	ExternalID 	  string 	 `json:"externalId"`                         
	ExpiredAt     *time.Time `json:"expiredAt"`
	PaidAt        *time.Time `json:"paidAt"`                       
}

