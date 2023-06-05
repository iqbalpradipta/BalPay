package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/iqbalpradipta/BalPay/features/user"
)

type Core struct {
	ID			uuid.UUID
	Balance		uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
	User		user.Core		
}