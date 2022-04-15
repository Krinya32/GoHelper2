package main

import (
	"fmt"
	"log"
	"net/http"
)

//w - responceWriter (куда писать ответ)
//r - request (Откуда брать запрос)
// Обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi! I'm new web-server")
}

// Товарищ, который выбирает нужный обработчик в зависимости от адреса, по которому пришел запрос
func RequestHandler() {
	http.HandleFunc("/", GetGreet)               // Если придет запрос по адресу "/" то вызывай GetGreet
	log.Fatal(http.ListenAndServe(":8080", nil)) // запускаем веб сервер в режиме "слушания"
}
func main() {
	RequestHandler()
}
