package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/shuheishintani/quote-memo-api/src/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomStringNumber(n int) string {
	var sb strings.Builder
	number := "1234567890"
	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomUser() models.User {
	user := models.User{
		ID:              RandomString(10),
		Username:        RandomString(10),
		ProfileImageUrl: RandomString(10),
		Provider:        RandomString(10),
	}
	return user
}

func RandomBook() models.Book {
	book := models.Book{
		Title:         RandomString(10),
		ISBN:          RandomStringNumber(10),
		Author:        RandomString(10),
		Publisher:     RandomString(10),
		CoverImageUrl: RandomString(10),
	}
	return book
}

func RandomTag() models.Tag {
	tag := models.Tag{
		Name: RandomString(6),
	}
	return tag
}

func RandomQuote(uid string, published bool) models.Quote {
	quote := models.Quote{
		Text:      RandomString(10),
		Page:      RandomInt(1, 500),
		Published: published,
		Book:      RandomBook(),
		Tags:      []models.Tag{RandomTag(), RandomTag(), RandomTag()},
		UserID:    uid,
	}
	return quote
}

func IncompleteRandomQuote(uid string, published bool, book models.Book, tags []models.Tag) models.Quote {
	quote := models.Quote{
		Text:      RandomString(10),
		Page:      RandomInt(1, 500),
		Published: published,
		BookID:    book.ID,
		Book:      book,
		Tags:      tags,
		UserID:    uid,
	}
	return quote
}
