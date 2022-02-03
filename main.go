package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Books model
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author `json:"author"`
}

type Author struct	{
	Firstname	string	`json:"firstname"`
	Lasttname	string	`json:"lastname"`
	AuthorID	string	`json:"authorid"`
}

//Init the book variable as a slice of the Book class
var books []Book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get a book with provided ID
func getBook(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get route params

	//shuffles through the books and find the matching item with provided detail
	for _, item := range books{
		if item.ID == params["id"]	{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//this tell the server that you are referencing from the Book struct
	json.NewEncoder(w).Encode(&Book{})
}

//Create a  book
func createBook(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update  a book with provided id
func updateBook(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if	item.ID == params["id"]	{
			//this makes a copy of book that macthes the provided ID by slicing
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete a book wtih provided id
func deleteBook(w http.ResponseWriter, r *http.Request)	{
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if	item.ID == params["id"]	{
			//this line deletes the particular book that macthes the provided ID by slicing
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main()	{
	//initialize router
	router := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "42342", Title: "Book One", Author: &Author{Firstname: "John", Lasttname: "Doe", AuthorID: "21655776"}})
	books = append(books, Book{ID: "2", Isbn: "41352", Title: "Book Two", Author: &Author{Firstname: "Steve", Lasttname: "Smith", AuthorID: "21635746"}})
	books = append(books, Book{ID: "3", Isbn: "41322", Title: "Book Three", Author: &Author{Firstname: "Mike", Lasttname: "Reagan", AuthorID: "21643978"}})

	//define the handler for endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}