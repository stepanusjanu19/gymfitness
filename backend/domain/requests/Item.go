package requests

type SaveItem struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateItem struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}