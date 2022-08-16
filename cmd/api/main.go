package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osalomon89/go-bookstore/internal/infrastructure/server"
)

func main() {
	httpRouter := mux.NewRouter()
	httpRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	}).Methods("GET")

	server.RegisterRouter(httpRouter)

	if err := http.ListenAndServe(":8080", httpRouter); err != nil {
		panic(err.Error())
	}
}
