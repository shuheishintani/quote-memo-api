package test

import (
	"log"
	"testing"

	"github.com/shuheishintani/quote-manager-api/src/services"
	"github.com/stretchr/testify/assert"
)

// func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		return nil, mock, err
// 	}

// 	gormDB, err := gorm.Open(
// 		postgres.New(
// 			postgres.Config{
// 				Conn: db,
// 			}), &gorm.Config{})

// 	if err != nil {
// 		return gormDB, mock, err
// 	}

// 	return gormDB, mock, err
// }

func TestGetTags(t *testing.T) {
	db, err := GormConnectForTesting()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	s := services.NewService(db)
	tags, err := s.GetTags()
	assert.NoError(t, err)
	assert.NotNil(t, tags)
}
