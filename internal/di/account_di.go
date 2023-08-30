package di

import (
	"github.com/gaogao-asia/golang-template/internal/account"
	"gorm.io/gorm"
)

func InitAccountHandler(db *gorm.DB) *account.AccountHandler {
	srv := account.NewAccountService(db)
	return account.NewAccountHandler(srv)
}
