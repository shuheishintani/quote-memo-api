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
	// db.Migrator().DropTable(&models.Quote{})
	// db.Migrator().DropTable(&models.Book{})
	// db.Migrator().DropTable(&models.Tag{})
	// db.Migrator().DropTable("quote_tags")
	db.AutoMigrate(&models.User{}, &models.Quote{}, &models.Book{}, &models.Tag{})

	return db, nil
}
