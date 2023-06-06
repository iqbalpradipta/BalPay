package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "github.com/iqbalpradipta/BalPay/features/user/data"
	deliveryUser "github.com/iqbalpradipta/BalPay/features/user/delivery"
	useCaseUser "github.com/iqbalpradipta/BalPay/features/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB)  {
	userDataFactory := userData.New(db)
	userUsecaseFactory := useCaseUser.NewLogic(userDataFactory)
	deliveryUser.NewHandler(e, userUsecaseFactory)
}