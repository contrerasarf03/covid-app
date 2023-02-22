package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Log        LogConfig        `mapstructure:"log"`
		API        ApiConfig        `mapstructure:"api"`
		Version    string           `mapstructure:"version"`
		Datasource DatasourceConfig `mapstructure:"datasource"`
	}

	ApiConfig struct {
		REST RestConfig `mapstructure:"rest"`
	}

	RestConfig struct {
		Host string     `mapstructure:"host"`
		Port int        `mapstructure:"port"`
		Spec string     `mapstructure:"spec"`
		CORS CorsConfig `mapstructure:"cors"`
	}

	CorsConfig struct {
		AllowedOrigins []string `mapstructure:"allowedOrigins"`
		AllowedHeaders []string `mapstructure:"allowedHeaders"`
		AllowedMethods []string `mapstructure:"allowedMethods"`
	}

	LogConfig struct {
		Level string `mapstructure:"level"`
	}

	DatasourceConfig struct {
		Type       string `mapstructure:"type"`
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		Database   string `mapstructure:"database"`
		Username   string `mapstructure:"db_username"`
		Password   string `mapstructure:"db_password"`
		SSLMode    string `mapstructure:"sslMode"`
		Migrations string `mapstructure:"migrations"`
	}
)

func defaultConfig() AppConfig {
	return AppConfig{
		Log: LogConfig{
			Level: "info",
		},
		API: ApiConfig{
			REST: RestConfig{
				Host: "0.0.0.0",
				Port: 3001,
				Spec: "./openapi.yaml",
				CORS: CorsConfig{
					AllowedOrigins: []string{"*"},
					AllowedHeaders: []string{
						"Content-Type",
						"Sec-Fetch-Dest",
						"Referer",
						"accept",
						"Sec-Fetch-Mode",
						"Sec-Fetch-Site",
						"User-Agent",
						"User-Agent",
						"API-KEY",
						"Authorization",
					},
					AllowedMethods: []string{
						"OPTIONS",
						"GET",
						"POST",
						"DELETE",
					},
				},
			},
		},
		Datasource: DatasourceConfig{
			Type:       "postgres",
			Host:       "localhost",
			Port:       5432,
			Database:   "postgres",
			Username:   "postgres",
			Password:   "postgres",
			SSLMode:    "disable",
			Migrations: "db/migrations",
		},
	}
}

func loadConfig() (config AppConfig, err error) {
	log.Info("Loading App Config...")

	config = defaultConfig()

	// --start--

	// parse rest config
	apiConfig := RestConfig{}
	err = viper.UnmarshalKey("api.rest", &apiConfig)
	if err != nil {
		log.WithError(err).Error("unable to read Rest API Config")
		os.Exit(1)
	}

	// parse log config
	logConfig := LogConfig{}
	err = viper.UnmarshalKey("log", &logConfig)
	if err != nil {
		log.WithError(err).Error("unable to read Log Config ")
		os.Exit(1)
	}

	// parse db config
	dbConfig := DatasourceConfig{}
	err = viper.UnmarshalKey("datasource", &dbConfig)
	if err != nil {
		log.WithError(err).Error("unable to read Datasource Config")
		os.Exit(1)
	} else {
		config.Datasource = dbConfig
	}

	config = AppConfig{
		API: ApiConfig{
			REST: apiConfig,
		},
		Log:        logConfig,
		Datasource: dbConfig,
	}
	// --end--

	return
}
