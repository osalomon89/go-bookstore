package server

import (
	"encoding/json"
	"net/http"

	"github.com/osalomon89/go-bookstore/internal/core/domain"
	"github.com/osalomon89/go-bookstore/internal/core/ports"
)

type AuthorController interface {
	Add(res http.ResponseWriter, req *http.Request)
}

type authorController struct {
	authorInteractor ports.AuthorService
}

func NewAuthorController(authorInteractor ports.AuthorService) AuthorController {
	return &authorController{authorInteractor}
}

func (controller *authorController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var author domain.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.authorInteractor.CreateAuthor(author)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
