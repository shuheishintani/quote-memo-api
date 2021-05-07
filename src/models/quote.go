package models

import "time"

type Quote struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Page      int       `json:"page"`
	ISBN      string    `json:"isbn"`
	Tags      []Tag     `gorm:"many2many:quote_tags;" json:"tags"`
}
