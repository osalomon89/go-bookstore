package usecases

import (
	"fmt"

	inventory "github.com/osalomon89/go-bookstore/inventory/domain"
)

type UseCasesInteractor interface {
	SaveBook(book inventory.Book) (*inventory.Book, error)
	ListInventory() ([]inventory.Book, error)
}

type useCasesInteractor struct {
	handler inventory.InventoryRepository
}

func NewUseCasesInteractor(handler inventory.InventoryRepository) UseCasesInteractor {
	return &useCasesInteractor{handler}
}

func (u *useCasesInteractor) SaveBook(book inventory.Book) (*inventory.Book, error) {
	err := u.handler.SaveBook(&book)
	if err != nil {
		return nil, fmt.Errorf("handler error: %w", err)
	}

	return &book, nil
}

func (u *useCasesInteractor) ListInventory() ([]inventory.Book, error) {
	results, _ := u.handler.ListInventory()
	return results, nil
}
