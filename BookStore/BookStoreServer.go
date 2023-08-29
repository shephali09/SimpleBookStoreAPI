package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BookStore struct {
	BookId    int    `json:"id"`
	BookName  string `json:"name"`
	Author    string `json:"AutName"`
	BookPrice int    `json:"price"`
}

var BookDetails = []BookStore{}

func addData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newBookDetails BookStore
		err := json.NewDecoder(r.Body).Decode(&newBookDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		BookDetails = append(BookDetails, newBookDetails)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Book Details added successfully!")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "Application/json")
		err := json.NewEncoder(w).Encode(BookDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func putData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updatedBook BookStore
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, book := range BookDetails {
			if book.BookId == updatedBook.BookId {
				BookDetails[i] = updatedBook
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Book updated successfully")
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		bookIDStr := r.URL.Query().Get("id")
		bookID, err := strconv.Atoi(bookIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, book := range BookDetails {
			if book.BookId == bookID {
				BookDetails = append(BookDetails[:i], BookDetails[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Book deleted Successfully")
			}

		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func main() {
	http.HandleFunc("/addBook", addData)
	http.HandleFunc("/getData", getData)
	http.HandleFunc("/updateBook", putData)
	http.HandleFunc("/deleteBook", deleteBook)

	http.ListenAndServe(":8081", nil)

}
