package main 

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

type Book struct {
	ID string `json:"id"`


	Isbn string `json:"isbn"`

	Title string `json:"title"`

	Author *Author `json:"author"`
}

type Author struct {

	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book


func getBooks(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
		json.NewEncoder(w).Encode(item)
		return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // id just for example - NOt safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
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

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main()  {

	r := mux.NewRouter()

	books = append (books, Book {ID: "1" , Isbn: "3434343" , Title:"Book one", Author: &Author {Firstname: "Mohammed", Lastname: "Aldawsari"} })

	books = append (books, Book {ID: "2" , Isbn: "32444444444" , Title:"Book Two", Author: &Author {Firstname: "Nasser", Lastname: "Steve"} })

	r.HandleFunc("/api/books", getBooks).Methods("GET")

	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	r.HandleFunc("/api/books", createBook).Methods("POST")

	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")

	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))


 

}