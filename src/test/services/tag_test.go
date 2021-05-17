package test

import (
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/shuheishintani/quote-memo-api/src/models"
	"github.com/shuheishintani/quote-memo-api/src/services"
	"github.com/shuheishintani/quote-memo-api/src/util"
	"github.com/stretchr/testify/assert"
)

func TestGetTags(t *testing.T) {
	db, err := gormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect gorm database: ", err)
	}

	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	defer postgresDB.Close()

	tag1 := util.RandomTag()
	tag2 := util.RandomTag()
	tag3 := util.RandomTag()
	tags := []models.Tag{
		tag1, tag2, tag3,
	}
	db.Create(&tags)

	s := services.NewService(db)
	result, err := s.GetTags()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(tags), len(result))
	assert.Equal(t, tag1.Name, result[0].Name)
	assert.Equal(t, tag2.Name, result[1].Name)
	assert.Equal(t, tag3.Name, result[2].Name)

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
}
