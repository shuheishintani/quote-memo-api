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

func TestPostQuote(t *testing.T) {
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

	postQuoteInput := models.Quote{
		Text:      util.RandomString(10),
		Page:      util.RandomInt(1, 500),
		Published: util.RandomBool(),
		Book:      book,
		Tags:      tags,
	}

	result, err := s.PostQuote(postQuoteInput, user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, postQuoteInput.Text, result.Text)
	assert.Equal(t, postQuoteInput.Page, result.Page)
	assert.Equal(t, postQuoteInput.Published, result.Published)
	assert.Equal(t, postQuoteInput.Book.Title, result.Book.Title)
	assert.Equal(t, postQuoteInput.Book.Author, result.Book.Author)
	assert.Equal(t, len(postQuoteInput.Tags), len(result.Tags))
	assert.Equal(t, len(tag1.Name), len(result.Tags[0].Name))
	assert.Equal(t, len(tag2.Name), len(result.Tags[1].Name))
	assert.Equal(t, len(tag3.Name), len(result.Tags[2].Name))

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetPublicQuotes(t *testing.T) {
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

	user1 := util.RandomUser()
	user2 := util.RandomUser()
	users := []models.User{
		user1, user2,
	}
	db.Create(&users)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	book := util.RandomBook()
	db.Create(&book)

	quote1 := util.IncompleteRandomQuote(user1.ID, true, book, tags)
	quote2 := util.RandomQuote(user1.ID, true)
	quote3 := util.RandomQuote(user1.ID, true)
	quote4 := util.RandomQuote(user1.ID, false)
	quote5 := util.RandomQuote(user2.ID, true)
	quote6 := util.RandomQuote(user2.ID, false)
	quotes := []models.Quote{
		quote1, quote2, quote3, quote4, quote5, quote6,
	}
	db.Create(&quotes)

	t.Run("Get public quotes", func(t *testing.T) {
		result, err := s.GetPublicQuotes([]string{}, 0, 10)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 4, len(result))
		assert.Equal(t, quotes[0].ID, result[0].ID)
		assert.Equal(t, quotes[0].Text, result[0].Text)
		assert.Equal(t, quotes[0].Page, result[0].Page)
		assert.Equal(t, quotes[0].Published, result[0].Published)
		assert.Equal(t, quotes[0].Book.Title, result[0].Book.Title)
		assert.Equal(t, quotes[0].Tags[0].Name, result[0].Tags[0].Name)
		assert.Equal(t, quotes[0].Tags[1].Name, result[0].Tags[1].Name)
	})

	t.Run("Get public quotes by specified tags", func(t *testing.T) {
		result, err := s.GetPublicQuotes([]string{tag1.Name, tag2.Name}, 0, 5)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, quotes[0].ID, result[0].ID)
		assert.Equal(t, quotes[0].Text, result[0].Text)
		assert.Equal(t, quotes[0].Page, result[0].Page)
		assert.Equal(t, quotes[0].Published, result[0].Published)
		assert.Equal(t, quotes[0].Book.Title, result[0].Book.Title)
		assert.Equal(t, quotes[0].Tags[0].Name, result[0].Tags[0].Name)
		assert.Equal(t, quotes[0].Tags[1].Name, result[0].Tags[1].Name)
	})

	t.Run("Limit and Offset are correctly working", func(t *testing.T) {
		result, err := s.GetPublicQuotes([]string{}, 2, 2)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, quotes[2].ID, result[0].ID)
		assert.Equal(t, quotes[2].Text, result[0].Text)
		assert.Equal(t, quotes[2].Page, result[0].Page)
		assert.Equal(t, quotes[2].Published, result[0].Published)
		assert.Equal(t, quotes[4].ID, result[1].ID)
		assert.Equal(t, quotes[4].Text, result[1].Text)
		assert.Equal(t, quotes[4].Page, result[1].Page)
		assert.Equal(t, quotes[4].Published, result[1].Published)
	})

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetQuotes(t *testing.T) {
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

	user1 := util.RandomUser()
	user2 := util.RandomUser()
	users := []models.User{
		user1, user2,
	}
	db.Create(&users)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	book := util.RandomBook()
	db.Create(&book)

	quote1 := util.IncompleteRandomQuote(user1.ID, false, book, tags)
	quote2 := util.RandomQuote(user1.ID, true)
	quote3 := util.RandomQuote(user1.ID, true)
	quote4 := util.RandomQuote(user1.ID, false)
	quotes := []models.Quote{
		quote1, quote2, quote3, quote4,
	}
	db.Create(&quotes)

	t.Run("Get private quotes", func(t *testing.T) {
		result, err := s.GetPrivateQuotes([]string{}, user1.ID, 0, 5)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 4, len(result))
		assert.Equal(t, quotes[0].ID, result[0].ID)
		assert.Equal(t, quotes[0].Text, result[0].Text)
		assert.Equal(t, quotes[0].Page, result[0].Page)
		assert.Equal(t, quotes[0].Published, result[0].Published)
		assert.Equal(t, quotes[0].Book.Title, result[0].Book.Title)
		assert.Equal(t, quotes[0].Tags[0].Name, result[0].Tags[0].Name)
		assert.Equal(t, quotes[0].Tags[1].Name, result[0].Tags[1].Name)
	})

	t.Run("Cannot get private quotes from other users", func(t *testing.T) {
		result, err := s.GetPrivateQuotes([]string{}, user2.ID, 0, 5)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, len(result))
	})

	t.Run("Get private quotes by specified tags", func(t *testing.T) {
		result, err := s.GetPrivateQuotes([]string{tag1.Name, tag2.Name}, user1.ID, 0, 5)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, quotes[0].ID, result[0].ID)
		assert.Equal(t, quotes[0].Text, result[0].Text)
		assert.Equal(t, quotes[0].Page, result[0].Page)
		assert.Equal(t, quotes[0].Published, result[0].Published)
		assert.Equal(t, quotes[0].Book.Title, result[0].Book.Title)
		assert.Equal(t, quotes[0].Tags[0].Name, result[0].Tags[0].Name)
		assert.Equal(t, quotes[0].Tags[1].Name, result[0].Tags[1].Name)
	})

	t.Run("Limit and Offset are correctly working", func(t *testing.T) {
		result, err := s.GetPrivateQuotes([]string{}, user1.ID, 2, 2)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, quotes[2].ID, result[0].ID)
		assert.Equal(t, quotes[2].Text, result[0].Text)
		assert.Equal(t, quotes[2].Page, result[0].Page)
		assert.Equal(t, quotes[2].Published, result[0].Published)
		assert.Equal(t, quotes[3].ID, result[1].ID)
		assert.Equal(t, quotes[3].Text, result[1].Text)
		assert.Equal(t, quotes[3].Page, result[1].Page)
		assert.Equal(t, quotes[3].Published, result[1].Published)
	})

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetFavoriteQuotes(t *testing.T) {
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

	user1 := util.RandomUser()
	user2 := util.RandomUser()
	users := []models.User{
		user1, user2,
	}
	db.Create(&users)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	book := util.RandomBook()
	db.Create(&book)

	quote1 := util.IncompleteRandomQuote(user1.ID, true, book, tags)
	quote2 := util.IncompleteRandomQuote(user1.ID, true, book, tags)

	quotes := []models.Quote{
		quote1, quote2,
	}
	db.Create(&quotes)

	db.Model(&user2).Association("FavoriteQuotes").Append(&quotes[0])
	db.Model(&user2).Association("FavoriteQuotes").Append(&quotes[1])

	result, err := s.GetFavoriteQuotes(user2.ID, 0, 10)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(quotes), len(result))
	assert.Equal(t, quotes[0].ID, result[0].ID)
	assert.Equal(t, quotes[0].Text, result[0].Text)
	assert.Equal(t, quotes[0].Page, result[0].Page)
	assert.Equal(t, quotes[0].Published, result[0].Published)
	assert.Equal(t, quotes[1].ID, result[1].ID)
	assert.Equal(t, quotes[1].Text, result[1].Text)
	assert.Equal(t, quotes[1].Page, result[1].Page)
	assert.Equal(t, quotes[1].Published, result[1].Published)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetQuoteById(t *testing.T) {
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

	user := util.RandomUser()
	db.Create(&user)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	book := util.RandomBook()
	db.Create(&book)

	quote1 := util.IncompleteRandomQuote(user.ID, false, book, tags)
	quote2 := util.RandomQuote(user.ID, true)
	quote3 := util.RandomQuote(user.ID, true)
	quote4 := util.RandomQuote(user.ID, false)
	quotes := []models.Quote{
		quote1, quote2, quote3, quote4,
	}
	db.Create(&quotes)

	strID := strconv.Itoa(quotes[0].ID)
	result, err := s.GetQuoteById(strID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, quotes[0].ID, result.ID)
	assert.Equal(t, quotes[0].Text, result.Text)
	assert.Equal(t, quotes[0].Page, result.Page)
	assert.Equal(t, quotes[0].Published, result.Published)
	assert.Equal(t, quotes[0].Book.Title, result.Book.Title)
	assert.Equal(t, quotes[0].Tags[0].Name, result.Tags[0].Name)
	assert.Equal(t, quotes[0].Tags[1].Name, result.Tags[1].Name)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestUpdateQuote(t *testing.T) {
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

	user := util.RandomUser()
	db.Create(&user)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	quote := util.RandomQuote(user.ID, true)
	db.Create(&quote)

	updateQuoteInput := models.Quote{
		Text:      util.RandomString(10),
		Page:      util.RandomInt(0, 500),
		Published: util.RandomBool(),
		Tags:      tags,
	}

	strID := strconv.Itoa(quote.ID)
	result, err := s.UpdateQuote(updateQuoteInput, strID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, quote.ID, result.ID)
	assert.Equal(t, updateQuoteInput.Text, result.Text)
	assert.Equal(t, updateQuoteInput.Page, result.Page)
	assert.Equal(t, updateQuoteInput.Published, result.Published)
	assert.Equal(t, updateQuoteInput.Tags[0].Name, result.Tags[0].Name)
	assert.Equal(t, updateQuoteInput.Tags[1].Name, result.Tags[1].Name)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestDeleteQuote(t *testing.T) {
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

	user := util.RandomUser()
	db.Create(&user)

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2,
	}
	db.Create(&tags)

	quote := util.RandomQuote(user.ID, true)
	db.Create(&quote)

	strID := strconv.Itoa(quote.ID)
	deleteResult, err := s.DeleteQuote(strID)

	notFoundResult := db.First(&models.Quote{}, quote.ID)

	assert.Error(t, notFoundResult.Error)
	assert.Equal(t, "record not found", notFoundResult.Error.Error())
	assert.NoError(t, err)
	assert.NotNil(t, deleteResult)
	assert.Equal(t, true, deleteResult)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestAddFavoriteQuote(t *testing.T) {
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

	user1 := util.RandomUser()
	user2 := util.RandomUser()
	users := []models.User{
		user1, user2,
	}
	db.Create(&users)

	quote := util.RandomQuote(users[0].ID, true)
	db.Create(&quote)

	strID := strconv.Itoa(quote.ID)
	result, err := s.AddFavoriteQuote(users[1].ID, strID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.FavoriteQuotes))
	assert.Equal(t, quote.ID, result.FavoriteQuotes[0].ID)
	assert.Equal(t, quote.Text, result.FavoriteQuotes[0].Text)
	assert.Equal(t, quote.Page, result.FavoriteQuotes[0].Page)
	assert.Equal(t, quote.Published, result.FavoriteQuotes[0].Published)
	assert.Equal(t, quote.Book.Title, result.FavoriteQuotes[0].Book.Title)
	assert.Equal(t, quote.Tags[0].Name, result.FavoriteQuotes[0].Tags[0].Name)
	assert.Equal(t, users[0].ID, result.FavoriteQuotes[0].UserID)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestRemoveFavoriteQuote(t *testing.T) {
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

	user1 := util.RandomUser()
	user2 := util.RandomUser()
	users := []models.User{
		user1, user2,
	}
	db.Create(&users)

	quote := util.RandomQuote(users[0].ID, true)
	db.Create(&quote)

	db.Model(&users[1]).Association("FavoriteQuotes").Append(&quote)

	strID := strconv.Itoa(quote.ID)
	result, err := s.RemoveFavoriteQuote(users[1].ID, strID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.FavoriteQuotes))

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}
