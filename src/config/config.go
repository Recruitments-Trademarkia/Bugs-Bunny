package config

import (
	"github.com/spf13/viper"
)

var (
	// DBConfig is the configuration for the database
	DBConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
	// ServerConfig is the configuration for the server
	ServerConfig struct {
		Port string
	}
)

// Init initializes the configuration
func Init() {
	DBConfig.Host = viper.GetString("DB_HOST")
	DBConfig.Port = viper.GetString("DB_PORT")
	DBConfig.Name = viper.GetString("DB_NAME")
	DBConfig.Username = viper.GetString("DB_USER")
	DBConfig.Password = viper.GetString("DB_PASS")
	ServerConfig.Port = viper.GetString("PORT")

}
