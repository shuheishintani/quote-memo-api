package test

import (
	"github.com/shuheishintani/quote-memo-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormConnectForTesting() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 port=5431 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&models.Quote{}, &models.Book{}, &models.Tag{}, &models.User{})
	return db, nil
}
