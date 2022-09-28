package entity

import "time"

type Book struct {
	ID          ID
	Title       string
	Description string
	AuthorID    ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//NewUser create a new user
func NewBook(title, description string, userId ID) *Book {
	book := &Book{
		ID:          NewID(),
		Title:       title,
		Description: description,
		AuthorID:    userId,
		CreatedAt:   time.Now(),
	}
	return book
}

//Validate validate data
func (u *Book) Validate() error {
	if u.Title == "" || u.Description == "" {
		return ErrInvalidEntity
	}

	return nil
}
