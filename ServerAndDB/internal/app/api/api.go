package api

import (
	"github.com/Krinya32/GoHelper2/ServerAndDB/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Base API server instance decsription
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}

//API constructor: build base API instance
func New(config *Config) *API {
	//log.Println("Cоздаем сервер с конфигурацией:")
	//log.Println("BinAddr: " + config.BindPort)
	//log.Println("LogLevel: " + config.LoggerLevel)
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure loggers,router,database connection and etc...
func (api *API) Start() error {
	// Trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//Подтверждение того что логгер сконфигурирован
	api.logger.Info("Starting api server at port", api.config.BindPort)

	//Конфигурируем маршрутизатор
	api.configureRouterField()
	//Конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}

	//На этапе валидного завершения стартуем http-сервер
	return http.ListenAndServe(api.config.BindPort, api.router)
}
