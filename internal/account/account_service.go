package account

import (
	"github.com/gaogao-asia/golang-template/internal/entity"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"gorm.io/gorm"
)

type accountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *accountService {
	return &accountService{
		db: db,
	}
}

func (s *accountService) GetAccounts() ([]entity.Account, error) {
	var accounts []entity.Account
	err := s.db.Find(&accounts).Error
	if err != nil {
		return nil, errs.ErrDBFailed.WithErr(err)
	}

	if len(accounts) == 0 {
		return nil, errs.ErrUserNotExist
	}

	return accounts, nil
}

func (s *accountService) CreateAccount(account *entity.Account) error {
	err := s.db.Create(&account).Error
	if err != nil {
		return errs.ErrDBFailed.WithErr(err)
	}

	return nil
}
