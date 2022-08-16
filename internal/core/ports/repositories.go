package ports

import "github.com/osalomon89/go-bookstore/internal/core/domain"

type BookRepository interface {
	SaveBook(book *domain.Book) error
	FindAll() (*[]domain.Book, error)
}

type AuthorRepository interface {
	SaveAuthor(author domain.Author) error
}
