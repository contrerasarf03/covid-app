package rest

import (
	"fmt"
	"net/http"

	covidController "github.com/Test/CovidApp/internal/entrypoint/controller"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type (
	API struct {
		config             *Config
		router             *mux.Router
		covidAppController covidController.Controller
	}
	Config struct {
		Host    string
		Port    int
		Spec    string
		Version string
		Cors    CORSConfig
	}

	CORSConfig struct {
		AllowedOrigins []string
		AllowedHeaders []string
		AllowedMethods []string
	}

	Database struct {
		DBHost  string
		DBUser  string
		DBName  string
		DBPass  string
		DBConns string
	}
)

func NewRestAPI(config *Config, router *mux.Router, covidAppController *covidController.Controller) *API {
	return &API{
		config:             config,
		router:             router,
		covidAppController: *covidAppController,
	}
}

func (api *API) Run() error {
	api.router = api.router.PathPrefix("/api/v1/covidapp").Subrouter()
	log.Info(api.config)
	api.registerHandlers()

	return http.ListenAndServe(api.address(), api.router)
}

func (api *API) address() string {
	return fmt.Sprintf("%s:%d", api.config.Host, api.config.Port)
}

func (api *API) registerHandlers() {
	api.covidAppController.RegisterHandlers(api.router)
}
