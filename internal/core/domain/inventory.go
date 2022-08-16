package domain

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name"`
	Age  int64  `json:"age"`
	City string `json:"city"`
}

type Book struct {
	gorm.Model
	Title    string
	Price    float64
	ISBN     string
	Stock    int
	AuthorID uint
}
