package books_api

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
	"net/http"
)

type Book struct {
    Id          string `json:"id"`
    ISBN        string `json:"isbn"`
    Title       string `json:"title"`
    Author      Author `json:"author"`
}

type Author struct {
    Firstname   string `json:"firstname"`
    Lastname    string `json:"lastname"`
}

type Message struct {
    Msg string `json:"msg"`
    Id  string `json:"id"`
}

var books []Book

func init() {
	PopulateBooks()
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range books {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(Message{"Book was not found", params["id"]})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    book := Book{}
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode("Invalid data")
        return
    }
    book.Id = uuid.New().String()
    books = append(books, book)
    json.NewEncoder(w).Encode(Message{"New book created", book.Id})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.Id == params["id"] {
            book := Book{}
            err := json.NewDecoder(r.Body).Decode(&book)
            if err != nil {
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode("Invalid data")
                return
            }
            books[index] = book
            msg := Message{"Book updated successfully", params["id"]}
            json.NewEncoder(w).Encode(msg)
            return
        }
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(Message{"Book was not found", params["id"]})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.Id == params["id"] {
            books = append(books[:index], books[index+1:]...)
            json.NewEncoder(w).Encode(Message{"Book deleted successfully",
                params["id"]})
            return
        }
    }
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(Message{"Book was not found", params["id"]})
}

func PopulateBooks() {
	    tmp_id := uuid.New().String()
    books = append(books, Book{
        tmp_id, "55432", "The Odissey", Author {"Luis", "cruz"}})
    tmp_id = uuid.New().String()
    books = append(books, Book{
        tmp_id, "394l0", "The Hobit", Author {"Andrea", "cruz"}})
    tmp_id = uuid.New().String()
    books = append(books, Book{
        tmp_id, "39099", "Da Vinci Code",Author {"Andrea", "Martinez"}})
}
