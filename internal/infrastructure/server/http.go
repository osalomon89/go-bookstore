package server

import (
	"github.com/gorilla/mux"
	"github.com/osalomon89/go-bookstore/internal/core/services"
	"github.com/osalomon89/go-bookstore/internal/infrastructure/db"
)

func RegisterRouter(router *mux.Router) {
	bookHandler, authorHandler := newHandlers()

	router.HandleFunc("/v1/books", bookHandler.FindAll).Methods("GET")
	router.HandleFunc("/v1/books", bookHandler.Add).Methods("POST")
	router.HandleFunc("/v1/authors", authorHandler.Add).Methods("POST")
}

func newHandlers() (BookController, AuthorController) {
	conn, err := db.GetConnectionDB()
	if err != nil {
		panic("error connecting to DB: " + err.Error())
	}

	err = db.Migrate(conn)
	if err != nil {
		panic("error running migrations: " + err.Error())
	}

	bookRepository, err := db.NewBookRepository(conn)
	if err != nil {
		panic("error creating book repository: " + err.Error())
	}

	authorRepository, err := db.NewAuthorRepository(conn)
	if err != nil {
		panic("error creating author repository: " + err.Error())
	}

	bookService := services.NewBookService(bookRepository)
	authorService := services.NewAuthorService(authorRepository)

	bookHandler := NewBookController(bookService)
	authorHandler := NewAuthorController(authorService)

	return bookHandler, authorHandler
}
