package main

import (
	"log"
	"net/http"
	"github.com/grabielcruz/books_app/books_api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

//	books_api.PopulateBooks()
	r.HandleFunc("/api/v1/books", books_api.GetBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", books_api.GetBook).Methods("GET")
	r.HandleFunc("/api/v1/books", books_api.CreateBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", books_api.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", books_api.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
