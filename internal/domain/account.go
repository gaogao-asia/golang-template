package domain

import (
	"context"
)

type Account struct {
	ID    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Roles string `gorm:"column:roles"` // seperate by comma: "admin,user"
}

//go:generate mockery --name AccountService --output ../../mocks
type AccountService interface {
	GetAccounts(ctx context.Context) ([]*Account, error)

	CreateAccount(ctx context.Context, acc *Account) error
}

//go:generate mockery --name AccountRepository --output ../../mocks
type AccountRepository interface {
	Get(ctx context.Context) ([]*Account, error)

	Create(ctx context.Context, acc *Account) error
}
