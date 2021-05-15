package models

import "time"

type Quote struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Page      int       `json:"page"`
	Published bool      `gorm:"default:false" json:"published"`
	Tags      []Tag     `gorm:"many2many:quote_tags;" json:"tags"`
	Book      Book      `json:"book"`
	BookID    int       `json:"book_id"`
	User      User      `json:"user"`
	UserID    string    `json:"user_id"`
}
