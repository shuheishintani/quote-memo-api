package test

import (
	"log"
	"testing"

	"github.com/shuheishintani/quote-manager-api/src/dto"
	"github.com/shuheishintani/quote-manager-api/src/models"
	"github.com/shuheishintani/quote-manager-api/src/services"
	"github.com/shuheishintani/quote-manager-api/src/util"
	"github.com/stretchr/testify/assert"
)

func TestGetQuotes(t *testing.T) {
	db, err := gormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect gorm database: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer postgresDB.Close()

	tags := []models.Tag{
		{Name: "tag1"},
		{Name: "tag2"},
		{Name: "tag3"},
	}
	db.Create(&tags)

	book := dto.Book{
		Title:         "book1",
		Isbn:          util.RandomStringNumber(10),
		Author:        "author1",
		Publisher:     util.RandomString(6),
		CoverImageUrl: util.RandomString(6),
	}
	db.Create(&book)

	postQuoteInput := dto.QuoteInput{
		Text: "quote1",
		Page: util.RandomInt(1, 500),
		Book: book,
		Tags: tags,
	}

	s := services.NewService(db)
	s.PostQuote(postQuoteInput, "randomUID")

	result, err := s.GetPrivateQuotes([]string{"tag1", "tag2", "tag3"}, "randomUID")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].ID, "randomUID")
	assert.Equal(t, result[0].Text, "quote1")

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
}

func TestPostQuote(t *testing.T) {
	db, err := gormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect gorm database: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer postgresDB.Close()

	book := dto.Book{
		Title:         util.RandomString(6),
		Isbn:          util.RandomStringNumber(10),
		Author:        util.RandomString(6),
		Publisher:     util.RandomString(6),
		CoverImageUrl: util.RandomString(6),
	}
	db.Create(&book)

	tags := []models.Tag{
		{Name: util.RandomString(6)},
		{Name: util.RandomString(6)},
		{Name: util.RandomString(6)},
	}
	db.Create(&tags)

	postQuoteInput := dto.QuoteInput{
		Text: util.RandomString(6),
		Page: util.RandomInt(1, 500),
		Book: book,
		Tags: tags,
	}

	s := services.NewService(db)
	result, err := s.PostQuote(postQuoteInput, util.RandomStringNumber(10))

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Text, postQuoteInput.Text)
	assert.Equal(t, len(result.Tags), len(postQuoteInput.Tags))
	assert.Equal(t, result.Book.Title, postQuoteInput.Book.Title)

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
}
