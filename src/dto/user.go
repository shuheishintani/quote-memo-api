package dto

type UserInput struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
	Provider        string `json:"provider"`
}
