package dto

import "github.com/shuheishintani/quote-manager-api/src/models"

type PostQuoteInput struct {
	Text string       `json:"text"`
	Page int          `json:"page"`
	ISBN string       `json:"isbn"`
	Tags []models.Tag `json:"tags"`
}
