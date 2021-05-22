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
	db, err := gormConnectForTesting()
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
		assert.Equal(t, updatedUser.Username, updateResult.Username)
		assert.Equal(t, updatedUser.ProfileImageUrl, updateResult.ProfileImageUrl)
		assert.Equal(t, updatedUser.Provider, updateResult.Provider)
		assert.Equal(t, updatedUser.Username, updateResult.Username)
	})

}
