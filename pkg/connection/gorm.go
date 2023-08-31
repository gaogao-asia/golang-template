package connection

import (
	"log"

	"github.com/gaogao-asia/golang-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Conn struct {
	DB *gorm.DB
}

func GetConnection() (Conn, error) {
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}
	url := config.AppConfig.Database.Postgres.GetDSN()
	pgDB, err := gorm.Open(postgres.Open(url), gormConfig)
	if err != nil {
		log.Println(err)
		return Conn{}, err
	}

	return Conn{
		DB: pgDB,
	}, nil
}
