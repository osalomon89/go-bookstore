package db

import (
	"fmt"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //nolint:gochecknoglobals

func GetConnectionDB() (*gorm.DB, error) {
	load()
	var err error

	if db == nil {
		db, err = gorm.Open(mysql.Open(dbConnectionURL()), &gorm.Config{})
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")

			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(domain.Author{}, domain.Book{})
	if err != nil {
		fmt.Printf("########## MIGRATE ERROR: " + err.Error() + " #############")

		return fmt.Errorf("### MIGRATE ERROR: %w", err)
	}

	return nil
}
