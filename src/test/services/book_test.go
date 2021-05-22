package test

import (
	"log"
	"strconv"
	"testing"

	"github.com/shuheishintani/quote-memo-api/src/models"
	"github.com/shuheishintani/quote-memo-api/src/services"
	"github.com/shuheishintani/quote-memo-api/src/util"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	db, err := GormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect gorm database: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer postgresDB.Close()

	s := services.NewService(db)

	book1 := util.RandomBook()
	book2 := util.RandomBook()
	book3 := util.RandomBook()
	books := []models.Book{
		book1, book2, book3,
	}
	db.Create(&books)

	result, err := s.GetBooks()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(books), len(result))
	assert.Equal(t, books[0].ID, result[0].ID)
	assert.Equal(t, books[0].Title, result[0].Title)
	assert.Equal(t, books[0].Author, result[0].Author)
	assert.Equal(t, books[0].ISBN, result[0].ISBN)
	assert.Equal(t, books[0].Publisher, result[0].Publisher)
	assert.Equal(t, books[0].CoverImageUrl, result[0].CoverImageUrl)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetBookById(t *testing.T) {
	db, err := GormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect gorm database: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer postgresDB.Close()

	s := services.NewService(db)

	book := util.RandomBook()
	db.Create(&book)

	strID := strconv.Itoa(book.ID)
	result, err := s.GetBookById(strID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, book.ID, result.ID)
	assert.Equal(t, book.Title, result.Title)
	assert.Equal(t, book.Author, result.Author)
	assert.Equal(t, book.ISBN, result.ISBN)
	assert.Equal(t, book.CoverImageUrl, result.CoverImageUrl)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}
