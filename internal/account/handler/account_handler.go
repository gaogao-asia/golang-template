package handler

import (
	"net/http"
	"strings"

	"github.com/gaogao-asia/golang-template/internal/account/dto"
	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/internal/server/http/response"
	"github.com/gaogao-asia/golang-template/pkg/requestbind"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountSrv domain.AccountService
}

func NewAccountHandler(accountSrv domain.AccountService) *AccountHandler {
	return &AccountHandler{
		accountSrv: accountSrv,
	}
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {
	accounts, err := h.accountSrv.GetAccounts(c.Request.Context())
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ResponseBody{
		Data: dto.GetAccountsResponse{
			Accounts: toGetAccountsResponse(accounts),
		},
	})
}

// toGetAccountsResponse
func toGetAccountsResponse(data []*domain.Account) []dto.AccountResponse {
	res := make([]dto.AccountResponse, 0)
	for _, v := range data {
		res = append(res, toAccountsResponse(v))
	}
	return res
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	req, err := requestbind.BindJson[dto.CreateAccountRequest](c)
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
	account := domain.Account{
		Name:  req.Name,
		Email: req.Email,
		Roles: strings.Join(req.Roles, ","),
	}
	err = h.accountSrv.CreateAccount(c.Request.Context(), &account)
	if err != nil {
		response.GeneralError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ResponseBody{
		Data: dto.CreateAccountResponse{
			Account: toAccountsResponse(&account),
		},
	})
}

func toAccountsResponse(data *domain.Account) dto.AccountResponse {
	res := dto.AccountResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Roles: func(roles string) []string {
			res := strings.Split(roles, ",")
			return res
		}(data.Roles),
	}
	return res
}
