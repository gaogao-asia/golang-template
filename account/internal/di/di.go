//go:build wireinject

package di

import (
	"github.com/gaogao-asia/golang-template/internal/account/handler"
	accrepo "github.com/gaogao-asia/golang-template/internal/account/repository"
	accsrv "github.com/gaogao-asia/golang-template/internal/account/service"
	"github.com/google/wire"
)

func InitAccountHandler() *handler.AccountHandler {
	panic(wire.Build(
		initDB,
		accrepo.NewAccountRepository,
		accsrv.NewAccountService,
		handler.NewAccountHandler,
	))
}
