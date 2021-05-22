package models

import "time"

type Quote struct {
	ID            int       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Text          string    `json:"text" validate:"required,min=1,max=400"`
	Page          int       `json:"page" validate:"gte=0,lte=50560"`
	Published     bool      `gorm:"default:false" json:"published"`
	Book          Book      `json:"book" validate:"required"`
	Tags          []Tag     `gorm:"many2many:quotes_tags;" json:"tags" validate:"required"`
	BookID        int       `json:"book_id"`
	User          User      `json:"user"`
	UserID        string    `json:"user_id"`
	FavoriteUsers []User    `gorm:"many2many:users_quotes;" json:"favorite_users"`
}
