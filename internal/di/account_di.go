package di

import (
	"github.com/gaogao-asia/golang-template/internal/account/handler"
	accrepo "github.com/gaogao-asia/golang-template/internal/account/repository"
	accsrv "github.com/gaogao-asia/golang-template/internal/account/service"
	"gorm.io/gorm"
)

func InitAccountHandler(db *gorm.DB) *handler.AccountHandler {
	repo := accrepo.NewAccountRepository(db)
	srv := accsrv.NewAccountService(repo)
	return handler.NewAccountHandler(srv)
}
