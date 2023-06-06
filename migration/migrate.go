package migration

import (
	"github.com/iqbalpradipta/BalPay/features/user/data"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB)  {
	db.AutoMigrate(
		new(data.User),
	)
}