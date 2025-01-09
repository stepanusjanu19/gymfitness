package requests

type Register struct {
	FullName string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=guest member trainer admin"`
}