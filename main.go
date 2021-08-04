package main

import (
	"big-integers-calculator/api"
	"big-integers-calculator/web/handlers"
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
	router.HandleFunc("/api", api.MuliplyData).Methods("POST")
	return router
}
