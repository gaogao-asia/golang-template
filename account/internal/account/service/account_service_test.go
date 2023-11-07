package service

import (
	"context"
	"testing"

	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/mocks"
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAccounts(t *testing.T) {
	tests := []struct {
		name     string
		aRepo    domain.AccountRepository
		expected []*domain.Account
		isError  assert.ErrorAssertionFunc
	}{
		{
			name: "success, get accounts",
			aRepo: func() domain.AccountRepository {
				mockARepo := mocks.NewAccountRepository(t)
				mockARepo.On("Get", context.Background()).Return([]*domain.Account{
					{
						ID:    1,
						Name:  "gaogao",
						Email: "trainer.minhtran@gmail.com",
						Roles: "admin,user",
					},
				}, nil)
				return mockARepo
			}(),
			expected: []*domain.Account{
				{
					ID:    1,
					Name:  "gaogao",
					Email: "trainer.minhtran@gmail.com",
					Roles: "admin,user",
				},
			},
			isError: assert.NoError,
		},
		{
			name: "error, get error",
			aRepo: func() domain.AccountRepository {
				mockARepo := mocks.NewAccountRepository(t)
				mockARepo.On("Get", context.Background()).Return(nil, errs.ErrDBFailed)
				return mockARepo
			}(),
			expected: nil,
			isError:  assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewAccountService(tt.aRepo)

			result, err := srv.GetAccounts(context.Background())
			tt.isError(t, err)
			assert.EqualValues(t, tt.expected, result)
		})
	}
}

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name         string
		aRepo        domain.AccountRepository
		requestParam *domain.Account
		isError      assert.ErrorAssertionFunc
	}{
		{
			name: "success, get account",
			requestParam: &domain.Account{
				Name:  "gaogao",
				Email: "trainer.minhtran@gmail.com",
				Roles: "admin,user",
			},
			aRepo: func() domain.AccountRepository {
				mockARepo := mocks.NewAccountRepository(t)
				mockARepo.On("Create", mock.Anything, &domain.Account{
					Name:  "gaogao",
					Email: "trainer.minhtran@gmail.com",
					Roles: "admin,user",
				}).Return(nil)
				return mockARepo
			}(),
			isError: assert.NoError,
		},
		{
			name: "error, get error",
			requestParam: &domain.Account{
				Name:  "gaogao",
				Email: "trainer.minhtran@gmail.com",
				Roles: "admin,user",
			},
			aRepo: func() domain.AccountRepository {
				mockARepo := mocks.NewAccountRepository(t)
				mockARepo.On("Create", mock.Anything, &domain.Account{
					Name:  "gaogao",
					Email: "trainer.minhtran@gmail.com",
					Roles: "admin,user",
				}).Return(errs.ErrDBFailed)
				return mockARepo
			}(),
			isError: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewAccountService(tt.aRepo)

			err := srv.CreateAccount(context.Background(), tt.requestParam)
			tt.isError(t, err)
		})
	}
}
