package migration

import "github.com/iqbalpradipta/BalPay/golang/src/config"

func Migration() {
	err := config.DB.AutoMigrate(

	)

	if err != nil {
		panic(err)
	}
}