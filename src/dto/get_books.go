package dto

type GetBooksInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Page   string `json:"page"`
}
