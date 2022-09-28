package presenter

import "goter.com.vn/server/entity"

//User data
type Book struct {
	ID          entity.ID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorID    entity.ID `json:"author_id"`
}
