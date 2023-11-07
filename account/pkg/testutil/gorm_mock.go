package testutil

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGORMMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// create dialector
	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	// open the database
	gormdb, err := gorm.Open(dialector, &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return gormdb, mock
}
