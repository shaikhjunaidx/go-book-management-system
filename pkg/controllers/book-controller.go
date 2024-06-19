package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/go-book-management-system/pkg/models"
	"github.com/shaikhjunaidx/go-book-management-system/pkg/utils"
)

var NewBook models.Book

func GetBook(responseWriter http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(response)
}

func GetBookById(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Parse Error")
	}

	bookDetails, _ := models.GetBookById(id)
	response, _ := json.Marshal(bookDetails)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(response)
}

func CreateBook(responseWriter http.ResponseWriter, request *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody[models.Book](request, newBook)
	aBook := newBook.CreateBook()
	response, _ := json.Marshal(aBook)

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(response)
}

func DeleteBook(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Parse Error")
	}

	book := models.DeleteBook(id)
	response, _ := json.Marshal(book)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(response)
}

func UpdateBook(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Parse Error")
	}

	var updateBook = &models.Book{}
	utils.ParseBody[models.Book](request, updateBook)

	book, db := models.GetBookById(id)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	response, _ := json.Marshal(book)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(response)
}
