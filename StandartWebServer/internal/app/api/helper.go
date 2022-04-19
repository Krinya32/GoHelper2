package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

//Пытаемся отконфигурировать наш API instance (а конкретнее поле Logger)
func (a *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(logLevel)
	return nil
}

//Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! this is rest api!"))

	})
}
