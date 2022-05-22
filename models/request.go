package models

type RequestModels struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
	Rating      int    `json:"rating" validate:"required,number"`
	Image       string `json:"image" validate:"required,max=255"`
}
