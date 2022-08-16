package ports

import "github.com/osalomon89/go-bookstore/internal/core/domain"

type AuthorService interface {
	CreateAuthor(author domain.Author) error
}

type BookService interface {
	CreateBook(book domain.Book) (*domain.Book, error)
	FindAll() (*[]domain.Book, error)
}
