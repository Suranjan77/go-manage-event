package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var P Properties

type Properties struct {
	Server     Server
	DataSource DataSource
}

type DataSource struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

type Server struct {
	Port string
}

func init() {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config-dev")
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if e, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Application configuration file not found")
		} else {
			panic("Error reading application configuration file\n" + e.Error())
		}
	}

	P = Properties{}

	err := viper.Unmarshal(&P)

	if err != nil {
		panic("Failed to load properties\n" + err.Error())
	}

	log.Println("Application properties loaded.")
}
