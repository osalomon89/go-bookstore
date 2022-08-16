package services

import (
	"log"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
)

type authorService struct {
	AuthorRepository ports.AuthorRepository
}

func NewAuthorService(AuthorRepository ports.AuthorRepository) ports.AuthorService {
	return &authorService{AuthorRepository}
}

func (interactor *authorService) CreateAuthor(author domain.Author) error {
	err := interactor.AuthorRepository.SaveAuthor(author)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
