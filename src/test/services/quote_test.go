package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/shuheishintani/quote-memo-api/src/dto"
	"github.com/shuheishintani/quote-memo-api/src/models"
	"github.com/shuheishintani/quote-memo-api/src/services"
	"github.com/shuheishintani/quote-memo-api/src/util"
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

	user := util.RandomUser()
	db.Create(&user)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tag3 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2, tag3,
	}
	db.Create(&tags)

	book := util.RandomBook()
	db.Create(&book)

	postQuoteInput := dto.QuoteInput{
		Text: util.RandomString(10),
		Page: util.RandomInt(1, 500),
		Book: book,
		Tags: tags,
	}

	s := services.NewService(db)
	quote1, err := s.PostQuote(postQuoteInput, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	result, err := s.GetPrivateQuotes([]string{tag1.Name, tag2.Name, tag3.Name}, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", result)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, quote1.ID, result[0].ID)
	assert.Equal(t, quote1.Text, result[0].Text)

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
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

	user := util.RandomUser()
	db.Create(&user)

	book := util.RandomBook()
	db.Create(&book)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tag3 := util.RandomTag()

	tags := []models.Tag{
		tag1, tag2, tag3,
	}
	db.Create(&tags)

	postQuoteInput := dto.QuoteInput{
		Text: util.RandomString(10),
		Page: util.RandomInt(1, 500),
		Book: book,
		Tags: tags,
	}

	s := services.NewService(db)
	result, err := s.PostQuote(postQuoteInput, user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, postQuoteInput.Text, result.Text)
	assert.Equal(t, postQuoteInput.Page, result.Page)
	assert.Equal(t, postQuoteInput.Book.Title, result.Book.Title)
	assert.Equal(t, len(postQuoteInput.Tags), len(result.Tags))
	assert.Equal(t, len(tag1.Name), len(result.Tags[0].Name))
	assert.Equal(t, len(tag2.Name), len(result.Tags[1].Name))
	assert.Equal(t, len(tag3.Name), len(result.Tags[2].Name))

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}
