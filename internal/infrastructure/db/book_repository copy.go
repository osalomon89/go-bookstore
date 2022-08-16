package db

import (
	"fmt"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
	"gorm.io/gorm"
)

type bookRepository struct {
	conn *gorm.DB
}

func NewBookRepository(conn *gorm.DB) (ports.BookRepository, error) {
	if conn == nil {
		return nil, fmt.Errorf("mysql connection cannot be nil")
	}

	return &bookRepository{
		conn: conn,
	}, nil
}

func (m *bookRepository) SaveBook(book *domain.Book) error {
	if err := m.conn.Create(book).Error; err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	return nil
}

func (m *bookRepository) FindAll() (*[]domain.Book, error) {
	books := new([]domain.Book)
	if err := db.Find(books).Error; err != nil {
		return nil, fmt.Errorf("error getting all books: %w", err)
	}

	return books, nil
}
