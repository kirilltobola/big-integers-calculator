package main

import (
	"big-integers-calculator/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.ListenAndServe(":8080", NewRouter())
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/", handlers.IndexGetHandler).Methods("GET")
	router.HandleFunc("/", handlers.IndexPostHandler).Methods("POST")
	return router
}
