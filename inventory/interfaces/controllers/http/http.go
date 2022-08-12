package http

import (
	"encoding/json"
	"net/http"

	inventory "github.com/osalomon89/go-bookstore/inventory/domain"
	usecases "github.com/osalomon89/go-bookstore/inventory/usecases"
)

type HTTPInteractor interface {
	Add(res http.ResponseWriter, req *http.Request)
	GetAll(res http.ResponseWriter, req *http.Request)
}

type httpInteractor struct {
	handler usecases.UseCasesInteractor
}

type ErrorResponse struct {
	Message string `json:"error"`
}

type BookRequestBody struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ISBN     string  `json:"isbn"`
	Stock    int     `json:"stock"`
	AuthorID uint    `json:"author_id"`
}

func NewHTTPInteractor(handler usecases.UseCasesInteractor) (HTTPInteractor, error) {
	return &httpInteractor{handler}, nil
}

func (h *httpInteractor) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book BookRequestBody

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	result, err := h.handler.SaveBook(inventory.Book{
		Title:    book.Title,
		Price:    book.Price,
		ISBN:     book.ISBN,
		Stock:    book.Stock,
		AuthorID: book.AuthorID,
	})

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	jsonResponse, jsonError := json.Marshal(result)
	if jsonError != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: jsonError.Error()})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}

func (h *httpInteractor) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	results, err := h.handler.ListInventory()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
