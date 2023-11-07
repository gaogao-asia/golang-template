package iaccount

import (
	"net/http"
	"testing"

	itest "github.com/gaogao-asia/golang-template/integration_test"

	adto "github.com/gaogao-asia/golang-template/internal/account/dto"
	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var (
	getAccountsURLPath = "/api/v1/accounts"
)

func TestGetAccounts(t *testing.T) {
	tests := []struct {
		name    string
		handler func(t *testing.T)
	}{
		{
			name:    getAccountsURLPath + " success",
			handler: getAccountSuccess,
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.handler)
	}
}

// getUserByPhoneSuccess
func getAccountSuccess(t *testing.T) {
	idata := itest.PrepareAccounts(t)
	defer idata.Cleanup()

	w, err := itest.CallAPI(http.MethodGet, getAccountsURLPath, nil)
	assert.NoError(t, err)

	response := w.Result()
	statusCode, res, err := itest.ParseResponse(response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, statusCode)

	var dataRes adto.GetAccountsResponse
	bytes, err := utils.JsonMarshal(res.Data)
	assert.NoError(t, err)
	err = utils.JsonUnmarshal(bytes, &dataRes)
	assert.NoError(t, err)

	// check len of accounts
	assert.Equal(t, len(idata.Accounts), len(dataRes.Accounts))

	// convert account response to map
	resAccount := make(map[int64]*domain.Account)
	lo.Map(idata.Accounts, func(v *domain.Account, i int) int {
		resAccount[v.ID] = v
		return i
	})

	// check account response
	for i, v := range idata.Accounts {
		if rv, ok := resAccount[v.ID]; ok {
			assert.Equal(t, v.ID, rv.ID)
			assert.Equal(t, v.Name, rv.Name)
			assert.Equal(t, v.Email, rv.Email)
			assert.Equal(t, v.Roles, rv.Roles)
		} else {
			t.Errorf("account %d not found", i)
		}
	}
}
