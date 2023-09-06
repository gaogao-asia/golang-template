package account

import (
	"net/http"

	"github.com/gaogao-asia/golang-template/internal/entity"
	"github.com/gaogao-asia/golang-template/internal/server/http/response"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountSrv *AccountService
}

func NewAccountHandler(accountSrv *AccountService) *AccountHandler {
	return &AccountHandler{
		accountSrv: accountSrv,
	}
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {
	accounts, err := h.accountSrv.GetAccounts()
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ResponseBody{
		Data: GetAccountsResponse{
			Accounts: toGetAccountsResponse(accounts),
		},
	})
}

// toGetAccountsResponse
func toGetAccountsResponse(data []entity.Account) []AccountResponse {
	res := make([]AccountResponse, 0)
	for _, v := range data {
		res = append(res, toAccountsResponse(v))
	}
	return res
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	req := CreateAccountRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	err = req.Validate()
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	// create account in database
	account := entity.Account{
		Name:  req.Name,
		Email: req.Email,
	}
	err = h.accountSrv.CreateAccount(&account)
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ResponseBody{
		Data: CreateAccountResponse{
			Account: toAccountsResponse(account),
		},
	})
}

func toAccountsResponse(data entity.Account) AccountResponse {
	res := AccountResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
	return res
}
