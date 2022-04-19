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
	log.Println(config.LoggerLevel + " созданный log level")
	log.Println(config.BindPort + " созданный BindArr")

	//_, err := toml.DecodeFile(configPath, config) // Десериализация содержимого .toml файла
	//if err != nil {
	//	log.Println("Can not find configs file, using default values:", err)
	//}

	parseTomlFile(config)

	log.Println(config.LoggerLevel + " спаршенный log level")
	log.Println(config.BindPort + " спаршенный parseBindArr")
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}

func parseTomlFile(config *api.Config) {
	_, err := toml.DecodeFile(configPath, config) // Десериализация содержимого .toml файла
	if err != nil {
		log.Println("Can not find configs file, using default values:", err)
	}
}
