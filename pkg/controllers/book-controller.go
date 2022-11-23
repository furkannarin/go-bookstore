package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	// "github.com/furkannarin/go-bookstore/pkg/utils"
	"github.com/furkannarin/go-bookstore/pkg/models"
	"github.com/furkannarin/go-bookstore/pkg/utils"
)

var NewBook models.Book

// json.Marshal, converts the argument it takes into a JSON file
// we need this because we work with http
func GetBook(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	NewBook := &models.Book{}
	utils.ParseJSON(request, &NewBook)
	book := models.CreateBook(NewBook)
	var res []byte
	res, _ = json.Marshal(&book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {}
func DeleteBook(response http.ResponseWriter, request *http.Request) {}
