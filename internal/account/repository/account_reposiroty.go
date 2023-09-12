package repository

import (
	"context"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		DB: db,
	}
}

func (r *accountRepository) Get(ctx context.Context) ([]*domain.Account, error) {
	var accounts []*domain.Account
	err := r.DB.Debug().WithContext(ctx).Find(&accounts).Error
	if err != nil {
		return nil, errs.ErrDBFailed.WithErr(err)
	}

	if len(accounts) == 0 {
		return nil, errs.ErrUserNotExist
	}
	return accounts, nil
}

func (r *accountRepository) Create(ctx context.Context, account *domain.Account) error {
	err := r.DB.Debug().Create(&account).Error
	if err != nil {
		return errs.ErrDBFailed.WithErr(err)
	}

	return nil
}
