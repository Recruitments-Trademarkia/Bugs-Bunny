package utils

import (
	"Bugs-Bunny/src/constants"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", 3000)
	viper.SetDefault("MIGRATE", false)
	viper.SetDefault("ENVIRONMENT", "development")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found ignoring error
		} else {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	for _, element := range constants.Env {
		if viper.GetString(element) == "" {
			log.Panicln(fmt.Errorf("env variables not present %s", element))
		}
	}

}
