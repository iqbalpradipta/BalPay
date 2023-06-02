package postgresql

import (
	"fmt"
	"log"

	"github.com/iqbalpradipta/BalPay/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to DB", err)
	}
	return db
}