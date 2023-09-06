package entity

type Account struct {
	ID    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Roles string `gorm:"column:roles"`
}

//go:generate mockery --name AccountService --output ../../mocks
type AccountService interface {
	GetAccounts() ([]Account, error)

	CreateAccount(account *Account) error
}
