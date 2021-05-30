package test

import (
	"log"
	"testing"

	"github.com/shuheishintani/quote-memo-api/src/models"
	"github.com/shuheishintani/quote-memo-api/src/services"
	"github.com/shuheishintani/quote-memo-api/src/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateUser(t *testing.T) {
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

	newUser := util.RandomUser()

	createResult, err := s.CreateOrUpdateUser(newUser)

	t.Run("Create user", func(t *testing.T) {

		assert.NoError(t, err)
		assert.NotNil(t, createResult)
		assert.Equal(t, newUser.Username, createResult.Username)
		assert.Equal(t, newUser.ProfileImageUrl, createResult.ProfileImageUrl)
		assert.Equal(t, newUser.Provider, createResult.Provider)
		assert.Equal(t, newUser.Username, createResult.Username)
	})

	t.Run("Update user", func(t *testing.T) {
		updatedUser := models.User{
			ID:              createResult.ID,
			Username:        util.RandomString(10),
			ProfileImageUrl: util.RandomString(10),
			Provider:        util.RandomString(10),
		}

		updateResult, err := s.CreateOrUpdateUser(updatedUser)
		assert.NoError(t, err)
		assert.NotNil(t, updateResult)
		assert.Equal(t, createResult.ID, updateResult.ID)
		assert.Equal(t, updatedUser.Username, updateResult.Username)
		assert.Equal(t, updatedUser.ProfileImageUrl, updateResult.ProfileImageUrl)
		assert.Equal(t, updatedUser.Provider, updateResult.Provider)
		assert.Equal(t, updatedUser.Username, updateResult.Username)
	})

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetUsers(t *testing.T) {
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
	user3 := util.RandomUser()
	users := []models.User{
		user1, user2, user3,
	}
	db.Create(&users)

	result, err := s.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(users), len(result))
	assert.Equal(t, users[0].ID, result[0].ID)
	assert.Equal(t, users[0].Username, result[0].Username)
	assert.Equal(t, users[0].ProfileImageUrl, result[0].ProfileImageUrl)
	assert.Equal(t, users[0].Provider, result[0].Provider)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetUserById(t *testing.T) {
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

	result, err := s.GetUserById(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.ProfileImageUrl, result.ProfileImageUrl)
	assert.Equal(t, user.Provider, result.Provider)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestDeleteUser(t *testing.T) {
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

	quote1 := util.RandomQuote(users[0].ID, true)
	quote2 := util.RandomQuote(users[1].ID, true)
	quotes := []models.Quote{
		quote1, quote2,
	}
	db.Create(&quotes)

	db.Model(&users[0]).Association("FavoriteQuotes").Append(&quotes[1])

	deleteResult, err := s.DeleteUser(user1.ID)
	notFoundResult := db.First(&models.User{}, "id = ?", users[0].ID)

	assert.NoError(t, err)
	assert.NotNil(t, deleteResult)
	assert.Equal(t, true, deleteResult)
	assert.Error(t, notFoundResult.Error)
	assert.Equal(t, "record not found", notFoundResult.Error.Error())

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}

func TestGetUserBooks(t *testing.T) {
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

	user := util.RandomUser()
	user.Books = books
	db.Create(&user)

	result, err := s.GetUserBooks(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(books), len(result))
	assert.Equal(t, books[0].Title, result[0].Title)
	assert.Equal(t, books[0].ISBN, result[0].ISBN)
	assert.Equal(t, books[0].Author, result[0].Author)
	assert.Equal(t, books[0].Publisher, result[0].Publisher)

	db.Migrator().DropTable("quotes_tags")
	db.Migrator().DropTable("users_quotes")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
	db.Migrator().DropTable("users")
}
