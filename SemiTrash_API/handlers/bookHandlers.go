package handlers

import (
	"SemiTrash_API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncased to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindsBookById(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ....")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "Provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookId int = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncased to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	oldBook, ok := models.FindsBookById(id)
	var newBook models.Book
	if !ok {
		log.Println("Book not found in database . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{
			"provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Нужно заменить oldBook на NewBook в DB!

	newBook.ID = oldBook.ID
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(201)
		json.NewEncoder(writer).Encode(newBook)
	}
	del := models.DeleteBookById(oldBook.ID)
	if !del {
		log.Println("I can`t delete book with ID :", id)
		writer.WriteHeader(500)
		msg := models.Message{Message: "Book with that ID does not delete in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	models.DB = append(models.DB, newBook)
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	book, ok := models.FindsBookById(id)
	if !ok {
		log.Println("Book not found in database . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Нужно удалить book из DB
	del := models.DeleteBookById(book.ID)
	if !del {
		log.Println("I can`t delete book with ID :", id)
		writer.WriteHeader(500)
		msg := models.Message{Message: "Book with that ID does not delete in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := models.Message{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}
