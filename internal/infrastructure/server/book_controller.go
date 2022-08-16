package server

import (
	"encoding/json"
	"net/http"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
)

type BookController interface {
	Add(res http.ResponseWriter, req *http.Request)
	FindAll(res http.ResponseWriter, req *http.Request)
}

type bookController struct {
	bookInteractor ports.BookService
}

func NewBookController(bookInteractor ports.BookService) BookController {
	return &bookController{bookInteractor}
}

func (controller *bookController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book domain.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	result, err2 := controller.bookInteractor.CreateBook(book)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func (controller *bookController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.bookInteractor.FindAll()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
