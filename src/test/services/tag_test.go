package test

import (
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/shuheishintani/quote-memo-api/src/services"
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

	createFixtures(db)

	s := services.NewService(db)
	result, err := s.GetTags()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 20)

	db.Migrator().DropTable("quote_tags")
	db.Migrator().DropTable("quotes")
	db.Migrator().DropTable("books")
	db.Migrator().DropTable("tags")
}
