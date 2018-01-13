package main

//  mux to route requests, mux stands for HTTP request multiplexer: which matches an incoming request to against a list of routes (registered)
// import packages
import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:id`
	Isbn   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:author`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init struct book slice: slice is dynamic array size
var books []Book

// handle HTTP requests
// Get all books
// parmas w and r, they're TYPES http.ResponseWriter and *http.Request respectively.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get one book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "1234", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "12345", Title: "Book Two", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	// API endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// port for server
	log.Fatal(http.ListenAndServe(":8000", r))
}
