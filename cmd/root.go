package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	root = &cobra.Command{
		Use:   "covid-app",
		Short: "The api for Covid APP",
	}
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	cobra.OnInitialize(initConfig)
	root.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is /covidapp/config/covid-app.yaml")
}

func initConfig() {

	if configFile != "" {
		// load config file if provided via args
		viper.SetConfigFile(configFile)
		log.Info("Load config file:", configFile)
	} else {
		// else try to load it from the tmp directory
		configFile = "/tmp/config/covid-app.yaml"
		viper.SetConfigFile(configFile)
	}

	// enable environment vars
	viper.SetEnvPrefix("COVIDAPP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println(err)
	}
}

// Execute is the entrypoint for the application
func Execute(version string) {
	root.Version = version
	if err := root.Execute(); err != nil {
		log.WithError(err).Error("Cannot execute root command")
		os.Exit(1)
	}
}
