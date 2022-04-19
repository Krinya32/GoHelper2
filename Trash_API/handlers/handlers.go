package handlers

import (
	"Trash_API/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) { // Прописываем Хедеры
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)                   // StatusCode для запроса
	json.NewEncoder(writer).Encode(models.DB) // Сериализация + запись в writer
}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(request) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := models.ErrorMessage{Message: "do not use ID not supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id #:", id)
	pizza, ok := models.FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := models.ErrorMessage{Message: "pizza with that id does not exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

func CreatePizza(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Creating new pizza")
	var pizza models.Pizza

	err := json.NewDecoder(request.Body).Decode(&pizza)
	if err != nil {
		msg := models.ErrorMessage{Message: "Provided json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var newPizza int = len(models.DB) + 1
	pizza.ID = newPizza
	models.DB = append(models.DB, pizza)
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(pizza)

}
