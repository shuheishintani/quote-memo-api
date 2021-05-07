package config

import (
	"fmt"
	"os"

	"github.com/shuheishintani/quote-manager-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnect() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&models.Quote{}, &models.Tag{})

	return db, nil
}
