package main

import (
	"Trash_API/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var port string = "8080"

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", handlers.GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizzas/{id}", handlers.GetPizzaById).Methods("GET")
	router.HandleFunc("/pizza", handlers.CreatePizza).Methods("POST")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}
