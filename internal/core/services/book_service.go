package services

import (
	"log"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
)

type bookService struct {
	BookRepository ports.BookRepository
}

func NewBookService(repository ports.BookRepository) ports.BookService {
	return &bookService{repository}
}

func (interactor *bookService) CreateBook(book domain.Book) (*domain.Book, error) {
	err := interactor.BookRepository.SaveBook(&book)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &book, nil
}

func (interactor *bookService) FindAll() (*[]domain.Book, error) {
	results, err := interactor.BookRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return results, nil
}
