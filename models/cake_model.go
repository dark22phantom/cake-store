package models

type CakeModel struct {
	ID          int    `json:"id"`
	Title       string `json:"string"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
