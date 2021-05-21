package dto

type UserInput struct {
	ID              string `json:"id" validate:"required"`
	Username        string `json:"username" validate:"required,max=100"`
	ProfileImageUrl string `json:"profile_image_url" validate:"url"`
	Provider        string `json:"provider" validate:"required,ma=100"`
}
