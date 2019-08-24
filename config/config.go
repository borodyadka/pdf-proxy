package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	Address  string
	LogLevel string
)

type specification struct {
	Address  string `envconfig:"ADDRESS" default:":8080"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`
}

func init() {
	godotenv.Load()

	var config specification
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	Address = config.Address
	LogLevel = config.LogLevel
}
