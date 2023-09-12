package iaccount

import (
	"net/http"
	"strings"
	"testing"

	itest "github.com/gaogao-asia/golang-template/integration_test"
	adto "github.com/gaogao-asia/golang-template/internal/account/dto"
	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	createAccountsURLPath = "/api/v1/accounts"
)

func TestCreateAccounts(t *testing.T) {
	tests := []struct {
		name    string
		handler func(t *testing.T)
	}{
		{
			name:    createAccountsURLPath + " success",
			handler: createAccountSuccess,
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.handler)
	}
}

// getUserByPhoneSuccess
func createAccountSuccess(t *testing.T) {
	requestBody := strings.NewReader(`{
		"name": "gaogao",
		"email": "gaogao@gmail.com",
		"roles": ["admin","user"]
	}`)

	w, err := itest.CallAPI(http.MethodPost, createAccountsURLPath, requestBody)
	assert.NoError(t, err)

	response := w.Result()
	statusCode, res, err := itest.ParseResponse(response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, statusCode)

	var dataRes adto.CreateAccountResponse
	bytes, err := utils.JsonMarshal(res.Data)
	assert.NoError(t, err)
	err = utils.JsonUnmarshal(bytes, &dataRes)
	assert.NoError(t, err)

	// check account response
	assert.Equal(t, "gaogao", dataRes.Account.Name)
	assert.Equal(t, "gaogao@gmail.com", dataRes.Account.Email)
	assert.Equal(t, []string{"admin", "user"}, dataRes.Account.Roles)

	// Clean data
	itest.Conn.DB.Model(&domain.Account{}).Where("id = ?", dataRes.Account.ID).Delete(&domain.Account{})
}
