package handler

import (
	"net/http"
	"testing"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/mocks"
	"github.com/gaogao-asia/golang-template/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	tests := []struct {
		name     string
		aService domain.AccountService
		expected string
		isError  assert.ErrorAssertionFunc
	}{
		{
			name: "Get list accounts",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)
				mockAsrv.On("GetAccounts").Return([]domain.Account{
					{
						ID:    1,
						Name:  "Minh",
						Email: "minhtran.dn.it@gmail.com",
						Roles: "admin",
					},
				}, nil)
				return mockAsrv
			}(),
			expected: `{"data":{"accounts":[{"id":1,"name":"Minh","email":"minhtran.dn.it@gmail.com","roles":["admin"]}]}}`,
			isError:  assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAccountHandler(tt.aService)

			ctx, resWriter := test.GetGinTestContext()
			handler.GetAccounts(ctx)

			assert.EqualValues(t, http.StatusOK, resWriter.Code)

			resBody := resWriter.Body.String()
			assert.EqualValues(t, tt.expected, resBody)
		})
	}
}
