package account

type CreateAccountRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"  binding:"required"`
}

type CreateAccountResponse struct {
	Account AccountResponse `json:"account"`
}

type AccountResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetAccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}
