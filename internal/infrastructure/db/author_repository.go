package db

import (
	"fmt"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
	"gorm.io/gorm"
)

type authorRepository struct {
	conn *gorm.DB
}

func NewAuthorRepository(conn *gorm.DB) (ports.AuthorRepository, error) {
	if conn == nil {
		return nil, fmt.Errorf("mysql connection cannot be nil")
	}

	return &authorRepository{
		conn: conn,
	}, nil
}

func (m *authorRepository) SaveAuthor(author domain.Author) error {
	if err := m.conn.Create(author).Error; err != nil {
		return fmt.Errorf("error saving author: %w", err)
	}

	return nil
}
