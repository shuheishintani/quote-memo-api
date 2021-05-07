package models

type Book struct {
	// ID            uint      `gorm:"primary_key" json:"id"`
	// CreatedAt     time.Time `json:"created_at"`
	Isbn          string `json:"isbn"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	CoverImageUrl string `json:"coverImageUrl"`
}
