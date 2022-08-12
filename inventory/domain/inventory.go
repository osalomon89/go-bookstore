package domain

import "gorm.io/gorm"

type InventoryRepository interface {
	SaveBook(book *Book) error
	ListInventory() ([]Book, error)
}

type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Book struct {
	gorm.Model
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ISBN     string  `json:"isbn"`
	Stock    int     `json:"stock"`
	AuthorID uint
}
