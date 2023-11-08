package dto

import (
	"context"

	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gaogao-asia/golang-template/pkg/tracing"
)

const (
	RoleAdmin  string = "admin"
	RoleUser   string = "user"
	RoleUnknow string = "unknow"
)

type CreateAccountRequest struct {
	Name  string   `json:"name" binding:"required"`
	Email string   `json:"email" binding:"required,email"`
	Roles []string `json:"roles" binding:"required"`
}

func (s CreateAccountRequest) Validate(ctx context.Context) error {
	ctx, span := tracing.Start(ctx, nil)
	defer span.End(ctx, nil)

	for _, v := range s.Roles {
		switch v {
		case RoleAdmin, RoleUser:
		default:
			return errs.ErrCreateNewAccountRequestRoleInvalid
		}
	}

	return nil
}

type CreateAccountResponse struct {
	Account AccountResponse `json:"account,omitempty"`
}

type AccountResponse struct {
	ID    int64    `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Email string   `json:"email,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type GetAccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}

type GetProductRequest struct {
	ProductID int `json:"product_id"`
}
