package repository

import (
	"context"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gaogao-asia/golang-template/pkg/log"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{
		DB: db,
	}
}

func (r *accountRepository) Get(ctx context.Context) (accounts []*domain.Account, err error) {
	ctx, span := tracing.Start(ctx, nil)
	defer span.End(ctx, log.Print{"accounts": &accounts})

	err = r.DB.Debug().WithContext(ctx).Find(&accounts).Error
	if err != nil {
		return nil, errs.ErrDBFailed.WithErr(err)
	}

	if len(accounts) == 0 {
		return nil, errs.ErrUserNotExist
	}
	return accounts, nil
}

func (r *accountRepository) Create(ctx context.Context, account *domain.Account) error {
	ctx, span := tracing.Start(ctx, log.Print{"account": &account})
	defer span.End(ctx, nil)

	err := r.DB.Debug().WithContext(ctx).Create(&account).Error
	if err != nil {
		return errs.ErrDBFailed.WithErr(err)
	}

	return nil
}
