package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gaogao-asia/golang-template/internal/domain"
	"github.com/gaogao-asia/golang-template/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAccountRepositoryGet(t *testing.T) {
	tests := []struct {
		name    string
		db      *gorm.DB
		want    []*domain.Account
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success, get accounts",
			db: func() *gorm.DB {
				db, mock := testutil.GetGORMMock()

				columns := []string{"id", "name", "email", "roles"}
				mock.ExpectQuery(`SELECT \* FROM \"accounts\"`).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow(1, "gaogao", "gao.gao@gmail.com", "admin,user").
						AddRow(2, "minh", "trainer.minhtran@gmail.com", "user"))

				return db
			}(),
			want: []*domain.Account{
				{
					ID:    1,
					Name:  "gaogao",
					Email: "gao.gao@gmail.com",
					Roles: "admin,user",
				},
				{
					ID:    2,
					Name:  "minh",
					Email: "trainer.minhtran@gmail.com",
					Roles: "user",
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success, no account found",
			db: func() *gorm.DB {
				db, mock := testutil.GetGORMMock()

				columns := []string{"id", "name", "email", "roles"}
				mock.ExpectQuery(`SELECT \* FROM \"accounts\"`).
					WillReturnRows(sqlmock.NewRows(columns))

				return db
			}(),
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "failed, get error from db",
			db: func() *gorm.DB {
				db, mock := testutil.GetGORMMock()

				mock.ExpectQuery(`SELECT \* FROM \"accounts\"`).
					WillReturnError(assert.AnError)

				return db
			}(),
			want:    nil,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewAccountRepository(tt.db)

			got, err := repo.Get(context.Background())

			tt.wantErr(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

// write unittest for Create
func TestAccountRepositoryCreate(t *testing.T) {
	tests := []struct {
		name    string
		input   *domain.Account
		db      *gorm.DB
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success, create account",
			input: &domain.Account{
				Name:  "gaogao",
				Email: "trainer.minhtran@gmail.com",
				Roles: "admin,user",
			},
			db: func() *gorm.DB {
				db, mock := testutil.GetGORMMock()

				mock.ExpectQuery(`^INSERT INTO "accounts" \("name","email","roles"\) VALUES \(\$1,\$2,\$3\) RETURNING "id"$`).
					WithArgs("gaogao", "trainer.minhtran@gmail.com", "admin,user").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				return db
			}(),

			wantErr: assert.NoError,
		},
		{
			name: "failed, create account",
			input: &domain.Account{
				Name:  "gaogao",
				Email: "trainer.minhtran@gmail.com",
				Roles: "admin,user",
			},
			db: func() *gorm.DB {
				db, mock := testutil.GetGORMMock()

				mock.ExpectQuery(`^INSERT INTO "accounts" \("name","email","roles"\) VALUES \(\$1,\$2,\$3\) RETURNING "id"$`).
					WithArgs("gaogao", "trainer.minhtran@gmail.com", "admin,user").
					WillReturnError(assert.AnError)

				return db
			}(),
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewAccountRepository(tt.db)

			err := repo.Create(context.Background(), tt.input)
			tt.wantErr(t, err)
		})
	}
}
