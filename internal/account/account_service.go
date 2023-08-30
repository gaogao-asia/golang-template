package account

import (
	"github.com/gaogao-asia/golang-template/internal/entity"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (s *AccountService) Getaccounts() ([]entity.Account, error) {
	var accounts []entity.Account
	err := s.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountService) CreateAccount(account *entity.Account) error {
	err := s.db.Create(&account).Error
	if err != nil {
		return err
	}
	return nil
}
