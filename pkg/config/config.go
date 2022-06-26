package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

var c config
var s secrets

type config struct {
	ServerPort                int
	DbHost                    string
	DbPort                    string
	DbName                    string
	DbMaxPoolSize             int
	DbMaxIdleConn             int
	DbMaxConnLifeTimeDuration time.Duration
	LogLevel                  string
}

type secrets struct {
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
}

// Load loads application properties
func Load() {
	env := os.Getenv("ENV")
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.BindEnv("ENV", "DB_USERNAME", "DB_PASSWORD")

	loadConfig()
	loadSecrets()
}

// DBConnectionUrl creates and returns database connection URL
func DBConnectionUrl() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		s.DbUsername,
		s.DbPassword,
		c.DbHost,
		c.DbPort,
		c.DbName,
	)
}

// LogLevel level of loggin for the application
func LogLevel() string {
	return c.LogLevel
}

// Port server port on which this application listens to
func Port() int {
	return c.ServerPort
}

/*
  MaxOpenConnections maximum DB connections to keep open at a time.
  This value should be less than the max_connections allowed by database.
*/
func MaxOpenConnections() int {
	return c.DbMaxPoolSize
}

/*
 MaxIdleConnections maximum DB connections to keep idle
 This value should be a fraction of the MaxConnections depending upon the expected load on the application.
*/
func MaxIdleConnections() int {
	return c.DbMaxIdleConn
}

/*
 MaxConnectionLifeTimeDuration maximnum allowed age for a connection
 Its better to to set this duration based on max idle connections,
 the higher the max idle connection the lower the max connection life
*/
func MaxConnectionLifeTimeDuration() time.Duration {
	return c.DbMaxConnLifeTimeDuration
}

// loadConfig loads Config from config-env.yaml file
func loadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if e, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Print(err)
			panic("Application configuration file not found")
		} else {
			panic("Error reading application configuration file\n" + e.Error())
		}
	}

	c = config{}
	err := viper.Unmarshal(&c)

	if err != nil {
		panic("Failed to load properties" + err.Error())
	}
}

// loadSecrets loads secrets from environment variables
func loadSecrets() {
	s = secrets{
		DbUsername: viper.GetString("DB_USERNAME"),
		DbPassword: viper.GetString("DB_PASSWORD"),
	}
}
