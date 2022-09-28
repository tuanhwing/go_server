package service

import (
	"goter.com.vn/server/dto"
	"goter.com.vn/server/entity"
	"goter.com.vn/server/repository"
)

type BookService interface {
	Create(book dto.BookCreateDTO, userId entity.ID) (entity.ID, error)
	FinByID(id entity.ID) (*entity.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Create(book dto.BookCreateDTO, userId entity.ID) (entity.ID, error) {
	e := entity.NewBook(book.Title, book.Description, userId)

	return service.bookRepository.Create(e)
}

func (service *bookService) FinByID(id entity.ID) (*entity.Book, error) {
	return service.bookRepository.FinByID(id)
}
