package handler

import (
	"context"
	"testing"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/mocks"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gaogao-asia/golang-template/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
				mockAsrv.On("GetAccounts", context.Background()).Return([]*domain.Account{
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
		{
			name: "get error",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)
				mockAsrv.On("GetAccounts", context.Background()).Return(nil, errs.ErrUserNotExist)
				return mockAsrv
			}(),
			expected: `{"error":{"code":404001,"msg_code":"USER_NOT_EXIST"}}`,
			isError:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAccountHandler(tt.aService)

			ctx, resWriter := testutil.GetGinTestContext()
			handler.GetAccounts(ctx)

			resBody := resWriter.Body.String()
			assert.EqualValues(t, tt.expected, resBody)
		})
	}
}

// TestCreateAccount
func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name        string
		aService    domain.AccountService
		requestBody string
		expected    string
		isError     assert.ErrorAssertionFunc
	}{
		{
			name: "Create account successful, get account",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)
				mockAsrv.On("CreateAccount", mock.Anything, &domain.Account{
					Name:  "Minh",
					Email: "trainer.minhtran@gmail.com",
					Roles: "admin",
				}).Return(nil)
				return mockAsrv
			}(),
			requestBody: `{"name":"Minh","email":"trainer.minhtran@gmail.com","roles":["admin"]}`,
			expected:    `{"data":{"account":{"name":"Minh","email":"trainer.minhtran@gmail.com","roles":["admin"]}}}`,
			isError:     assert.NoError,
		},
		{
			name: "Create account miss request field, get error",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)
				return mockAsrv
			}(),
			requestBody: `{"name":"Minh","email":"trainer.minhtran@gmail.com"}`,
			expected:    `{"error":{"code":400000,"msg_code":"BAD_REQUEST"}}`,
			isError:     assert.Error,
		},
		{
			name: "Create account wrong role, get error",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)

				return mockAsrv
			}(),
			requestBody: `{"name":"Minh","email":"trainer.minhtran@gmail.com","roles":["fwef"]}`,
			expected:    `{"error":{"code":400001,"msg_code":"CREATE_ACCOUNT_REQUEST_ROLE_INVALID"}}`,
			isError:     assert.Error,
		},
		{
			name: "Create account service return error, get error",
			aService: func() domain.AccountService {
				mockAsrv := mocks.NewAccountService(t)
				mockAsrv.On("CreateAccount", mock.Anything, &domain.Account{
					Name:  "Minh",
					Email: "trainer.minhtran@gmail.com",
					Roles: "admin",
				}).Return(errs.ErrBadRequest)
				return mockAsrv
			}(),
			requestBody: `{"name":"Minh","email":"trainer.minhtran@gmail.com","roles":["admin"]}`,
			expected:    `{"error":{"code":400000,"msg_code":"BAD_REQUEST"}}`,
			isError:     assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewAccountHandler(tt.aService)

			ctx, resWriter := testutil.GetGinTestContextWithBody(tt.requestBody)
			handler.CreateAccount(ctx)

			resBody := resWriter.Body.String()
			assert.EqualValues(t, tt.expected, resBody)
		})
	}
}
