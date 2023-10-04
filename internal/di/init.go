package di

import (
	"github.com/gaogao-asia/golang-template/pkg/connection"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	return connection.DB
}
