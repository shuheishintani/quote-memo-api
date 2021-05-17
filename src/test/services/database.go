package test

import (
	"github.com/shuheishintani/quote-memo-api/src/models"
	"github.com/shuheishintani/quote-memo-api/src/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormConnectForTesting() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 port=5431 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return db, err
	}
	db.AutoMigrate(&models.Quote{}, &models.Book{}, &models.Tag{})

	return db, nil
}

func createFixtures(db *gorm.DB) {
	tags := []models.Tag{}
	books := []models.Book{}

	for i := 0; i < 20; i++ {
		tags = append(tags, models.Tag{Name: util.RandomString(6)})
	}
	db.Create(&tags)

	for i := 0; i < 20; i++ {
		tags = append(tags, models.Tag{Name: util.RandomString(6)})
	}
	db.Create(&tags)

	for i := 0; i < 20; i++ {
		books = append(books, models.Book{
			ISBN:          util.RandomStringNumber(10),
			Title:         util.RandomString(6),
			Author:        util.RandomString(6),
			Publisher:     util.RandomString(6),
			CoverImageUrl: util.RandomString(6),
		})
	}
	db.Create(&books)
}
