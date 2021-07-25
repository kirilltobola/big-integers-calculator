package main

import (
	"big-integers-calculator/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexGetHandler).Methods("GET")
	router.HandleFunc("/", handlers.IndexPostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
