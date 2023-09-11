package service

import (
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

func (s *accountService) GetAccounts() ([]*domain.Account, error) {
	accounts, err := s.accountRepo.Get()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *accountService) CreateAccount(account *domain.Account) error {
	err := s.accountRepo.Create(account)
	if err != nil {
		return err
	}

	return nil
}
