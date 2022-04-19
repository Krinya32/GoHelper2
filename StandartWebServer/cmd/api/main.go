package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Krinya32/GoHelper2/StandartWebServer/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml")
}

func main() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("It's works")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериализация содержимого .toml файла
	if err != nil {
		log.Println("Can not find configs file, using default values:", err)
	}
	// Теперь тут надо попробовать прочитать из .toml/.env , так как там может быть новая информация
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())

}
