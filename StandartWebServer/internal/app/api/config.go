package api

//General instance for API server of REST application
type Config struct {
	//Port
	BindPort string `toml:"bind_addr"`
	//Logger level
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindPort:    ":8080",
		LoggerLevel: "debug",
	}
}
