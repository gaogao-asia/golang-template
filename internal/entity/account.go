package entity

type Account struct {
	ID    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Roles string `gorm:"column:roles"`
}
