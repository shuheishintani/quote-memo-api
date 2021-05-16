package dto

import (
	"github.com/shuheishintani/quote-memo-api/src/models"
)

type Book struct {
	Isbn          string `json:"isbn"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type QuoteInput struct {
	Text      string       `json:"text"`
	Page      int          `json:"page"`
	Published bool         `json:"published"`
	Book      Book         `json:"book"`
	Tags      []models.Tag `json:"tags"`
}
