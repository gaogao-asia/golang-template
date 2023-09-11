package account

import "github.com/gaogao-asia/golang-template/pkg/errs"

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleUser   Role = "user"
	RoleUnknow Role = "unknow"
)

type CreateAccountRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Roles []Role `json:"roles" binding:"required"`
}

func (s CreateAccountRequest) Validate() error {
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
