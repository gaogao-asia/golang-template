package service

import (
	"context"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
)

type accountService struct {
	accountRepo domain.AccountRepository
}

func NewAccountService(accountRepo domain.AccountRepository) domain.AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

func (s *accountService) GetAccounts(ctx context.Context) (accounts []*domain.Account, err error) {
	ctx, span := tracing.Start(ctx, nil)
	defer span.End(ctx, log.Print{"accounts": &accounts})

	accounts, err = s.accountRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *accountService) CreateAccount(ctx context.Context, account *domain.Account) error {
	err := s.accountRepo.Create(ctx, account)
	if err != nil {
		return err
	}

	return nil
}
