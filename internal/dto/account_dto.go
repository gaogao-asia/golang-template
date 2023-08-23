package dto

type GetAccountResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateAccountRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"  binding:"required"`
}

type CreateAccountResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
