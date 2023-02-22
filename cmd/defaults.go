package cmd

import (
	"github.com/spf13/viper"
)

var defaults = map[string]interface{}{
	"api.rest.host":                "0.0.0.0",
	"api.rest.port":                8080,
	"api.rest.spec":                "./openapi.yaml",
	"api.rest.cors.allowedOrigins": []string{"*"},
	"api.rest.cors.allowedHeaders": []string{
		"Content-Type",
		"Referer",
		"Accept",
		"User-Agent",
		"Sec-Fetch-Dest",
		"Sec-Fetch-Mode",
		"Sec-Fetch-Site",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
		"X-Requested-With",
		"RA-API-KEY",
		"SECRET-API-KEY",
	},
	"api.rest.cors.allowedMethods": []string{
		"OPTIONS",
		"GET",
		"POST",
	},

	"log.debug": true,

	// Default configs for Checkout.com
	"checkout.<config-key>": "<config-value>",
}

func init() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}
