package mysqldb

import (
	"fmt"

	"github.com/osalomon89/go-bookstore/inventory/domain"
	"github.com/osalomon89/go-bookstore/inventory/interfaces/repository"
	"gorm.io/gorm"
)

type mySqlDB struct {
	conn *gorm.DB
}

func NewMySqlDB() (repository.RepositoryInteractorRespository, error) {
	Load()
	conn, err := getConnectionDB()
	if err != nil {
		return nil, fmt.Errorf("error in mysql connection: %w", err)
	}

	if err := Migrate(conn); err != nil {
		return nil, fmt.Errorf("error in mysql migration: %w", err)
	}

	return &mySqlDB{
		conn: conn,
	}, nil
}

func (m *mySqlDB) SaveBook(book *domain.Book) error {
	if err := m.conn.Create(book).Error; err != nil {
		return fmt.Errorf("error saving book: %w", err)
	}

	return nil
}

func (m *mySqlDB) ListInventory() ([]domain.Book, error) {
	return nil, nil
}
