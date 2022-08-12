package repository

import (
	"github.com/osalomon89/go-bookstore/inventory/domain"
)

type RepositoryInteractorRespository interface {
	SaveBook(book *domain.Book) error
	ListInventory() ([]domain.Book, error)
}

type repositoryInteractor struct {
	handler RepositoryInteractorRespository
}

func NewRepositoryInteractor(handler RepositoryInteractorRespository) RepositoryInteractorRespository {
	return &repositoryInteractor{handler}
}

func (r *repositoryInteractor) SaveBook(book *domain.Book) error {
	return r.handler.SaveBook(book)
}

func (r *repositoryInteractor) ListInventory() ([]domain.Book, error) {
	results, _ := r.handler.ListInventory()
	return results, nil
}
