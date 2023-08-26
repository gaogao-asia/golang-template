package connection

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("err loading: %v", e)
	}
	dnsReal := fmt.Sprintf(dsn, "localhost", os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))
	fmt.Println(dnsReal)
	db, err := gorm.Open(postgres.Open(dnsReal), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
