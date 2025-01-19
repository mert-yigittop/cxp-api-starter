package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectToPostgresql() (*gorm.DB, error) {
	source := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println(source)
	DB, err := gorm.Open(postgres.Open(source), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
