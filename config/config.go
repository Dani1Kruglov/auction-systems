package config

import (
	"errors"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = "8080"
)

type Config struct {
	PostgresDatabaseConfig *PostgresDatabaseConfig
	HttpHost               string
	HttpPort               string
}

type PostgresDatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

func NewConfig(fileName string) (*Config, error) {

	err := godotenv.Load(fileName)
	if err != nil {
		return nil, errors.New("error loading .env file: " + err.Error())
	}

	appConfig := new(Config)
	appConfig.PostgresDatabaseConfig = &PostgresDatabaseConfig{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}
	appConfig.HttpHost = os.Getenv("HTTP_HOST")
	appConfig.HttpPort = os.Getenv("HTTP_PORT")

	if appConfig.HttpHost == "" {
		appConfig.HttpHost = defaultHost
	}
	if appConfig.HttpPort == "" {
		appConfig.HttpPort = defaultPort
	}

	if appConfig.PostgresDatabaseConfig != nil && emptyFields(*appConfig.PostgresDatabaseConfig) {
		return nil, errors.New("one or more fields in PostgresDatabaseConfig are empty")
	}

	return appConfig, nil
}

func emptyFields(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type() == reflect.TypeOf("") && field.String() == "" {
			return true
		}
	}

	return false
}
