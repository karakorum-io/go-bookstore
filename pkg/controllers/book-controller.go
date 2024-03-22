package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karakorum.io/karakorum-book-store/pkg/models"
	"github.com/karakorum.io/karakorum-book-store/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	utils.SendResponse(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Invalid Id")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	utils.SendResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()

	res, _ := json.Marshal(b)
	utils.SendResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Invalid Id")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	utils.SendResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Invalid Id")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	utils.SendResponse(w, res)
}
