package responses

import "github.com/gofrs/uuid"

type SignUpResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
	IsActive bool   `json:"is_active"`
}

type ItemResponse struct {
	ItemID      uuid.UUID `json:"item_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}