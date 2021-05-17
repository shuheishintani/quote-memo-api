package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/shuheishintani/quote-memo-api/src/dto"
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

func RandomUser() models.User {
	user := models.User{
		ID:              RandomString(10),
		Username:        RandomString(10),
		ProfileImageUrl: RandomString(10),
		Provider:        RandomString(10),
	}
	return user
}

func RandomBook() dto.Book {
	book := dto.Book{
		Title:         RandomString(10),
		Isbn:          RandomStringNumber(10),
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
