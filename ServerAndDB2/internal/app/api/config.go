package api

import "github.com/Krinya32/GoHelper2/ServerAndDB2/storage"

//General instance for API server of REST application
type Config struct {
	//Port
	BindPort string `toml:"bind_addr"`
	//Logger level
	LoggerLevel string `toml:"logger_level"`
	//Store configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindPort:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
