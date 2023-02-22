//go:build wireinject
// +build wireinject

package cmd

import (
	"os"

	covidAppService "github.com/Test/CovidApp/internal/component"
	rest "github.com/Test/CovidApp/internal/entrypoint"
	covidAppHandler "github.com/Test/CovidApp/internal/entrypoint/controller"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/Test/CovidApp/internal/infrastructure/postgres"
	repository "github.com/Test/CovidApp/internal/infrastructure/postgres/repository"
	"github.com/spf13/viper"
)

func createRestAPI() *rest.API {
	wire.Build(
		ProvideConfig,
		ProvideGormDB,
		ProvideDatasource,

		// Service constructors
		covidAppService.NewCovidApp,
		repository.NewGormRepository,
		ProvideServiceConfig,

		// Controller
		covidAppHandler.NewController,

		// Rest constructors
		mux.NewRouter,
		ProvideRestAPIConfig,
		rest.NewRestAPI,
	)

	return &rest.API{}
}

func createMigration() *postgres.Migration {
	wire.Build(
		ProvideConfig,
		ProvideDatasource,

		postgres.NewMigration,
	)
	return &postgres.Migration{}
}

func ProvideRestAPIConfig(config AppConfig) *rest.Config {

	var corsConfig rest.CORSConfig
	corsConfig.AllowedOrigins = config.API.REST.CORS.AllowedOrigins
	corsConfig.AllowedHeaders = config.API.REST.CORS.AllowedHeaders
	corsConfig.AllowedMethods = config.API.REST.CORS.AllowedMethods

	restConfig := &rest.Config{
		Host:    config.API.REST.Host,
		Port:    config.API.REST.Port,
		Spec:    config.API.REST.Spec,
		Cors:    corsConfig,
		Version: "local",
	}

	log.Info("========================================")
	log.Info("API Configuration")
	log.Info("========================================")
	log.Info("Host:    ", restConfig.Host)
	log.Info("Port:    ", restConfig.Port)
	log.Info("Spec:    ", restConfig.Spec)
	log.Info("Version: ", restConfig.Port)

	return restConfig
}

func ProvideConfig() AppConfig {
	config, err := loadConfig()
	if err != nil {
		log.WithError(err).Error("unable to unmarshal configuration")
		os.Exit(1)
	}

	return config
}

func ProvideServiceConfig(config AppConfig) covidAppService.Config {
	return covidAppService.Config{
		Host: config.API.REST.Host,
		Port: config.API.REST.Port,
		Spec: config.API.REST.Spec,
	}
}

func ProvideDatabaseConfig() *rest.Database {
	var config rest.Database
	err := viper.UnmarshalKey("api.rest", &config)
	if err != nil {
		log.WithError(err).Error("unable to read database config")
		os.Exit(1)
	}

	log.Info("========================================")
	log.Info("API Configuration")
	log.Info("========================================")
	log.Info("Host:  ", config.DBHost)
	log.Info("User:  ", config.DBUser)
	log.Info("Conns: ", config.DBConns)

	return &config
}

func ProvideDatasource(config AppConfig) *postgres.Datasource {
	return &postgres.Datasource{
		Type:       config.Datasource.Type,
		Host:       config.Datasource.Host,
		Port:       config.Datasource.Port,
		Database:   config.Datasource.Database,
		Username:   config.Datasource.Username,
		Password:   config.Datasource.Password,
		SSLMode:    config.Datasource.SSLMode,
		Migrations: config.Datasource.Migrations,
	}
}

func ProvideGormDB(datasource *postgres.Datasource) *gorm.DB {
	db, err := gorm.Open("postgres", datasource.AsPQString())
	if err != nil {
		log.WithError(err).Error("unable to get gorm db connection")
		os.Exit(1)
	}

	return db
}
