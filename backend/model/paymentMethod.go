package model

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name        string `json:"name" form:"name"`                         
	Code        string `json:"code" form:"code"`                        
	Provider    string `json:"provider" form:"provider"`                 
	Active      bool   `json:"active" form:"active"`                    
	IconUrl     string `json:"iconUrl" form:"iconUrl"`

	Transactions []Transaction `gorm:"foreignKey:PaymentMethodID"`
}