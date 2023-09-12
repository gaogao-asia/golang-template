package service

import (
	"context"

	"github.com/gaogao-asia/golang-template/internal/domain"
)

type accountService struct {
	accountRepo domain.AccountRepository
}

func NewAccountService(accountRepo domain.AccountRepository) *accountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

func (s *accountService) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	accounts, err := s.accountRepo.Get(ctx)
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
