package test

import (
	"fmt"

	"github.com/shuheishintani/quote-manager-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnectForTesting() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	dbUsername := "postgres"
	dbPassword := "postgres"
	dbHost := "localhost"
	dbTable := "postgres"
	dbPort := "5431"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&models.Quote{}, &models.Book{}, &models.Tag{})

	return db, nil
}
