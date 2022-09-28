package dto

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" bindding:"required"`
	Description string `json:"description" form:"description" bindding:"required"`
}
