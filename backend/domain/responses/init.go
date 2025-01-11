package responses

type SignUpResponse struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	IsActive bool `json:"is_active"`
}