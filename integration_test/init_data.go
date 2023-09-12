package integrationtest

import (
	"testing"

	"github.com/gaogao-asia/golang-template/internal/domain"
)

type ITestData struct {
	Accounts []*domain.Account
}

func PrepareAccounts(t *testing.T) ITestData {
	accs := []*domain.Account{
		{
			Name:  "gaogao",
			Email: "gaogao@gmail.com",
			Roles: "admin,user",
		},
		{
			Name:  "minh",
			Email: "trainer.minhtran@gmail.com",
			Roles: "user",
		},
	}
	err := Conn.DB.Create(&accs).Error
	if err != nil {
		t.Fatal(err)
	}

	return ITestData{Accounts: accs}
}

func (s *ITestData) Cleanup() {
	if len(s.Accounts) > 0 {
		Conn.DB.Delete(s.Accounts)
	}
}
