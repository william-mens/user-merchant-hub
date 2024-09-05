package config

import (
	"os"
)

//field names should be uppercase

type Configuration struct {
	DbName           string
	Host             string
	Password         string
	Port             string
	DbUser           string
	TursoDbUrl       string
	TursoDBAuthToken string
}

func Config() Configuration {
	configuration := Configuration{
		DbName:           os.Getenv("DB_NAME"),
		Host:             os.Getenv("DB_HOST"),
		DbUser:           os.Getenv("DB_USERNAME"),
		Password:         os.Getenv("DB_PASSWORD"),
		Port:             os.Getenv("DB_PORT"),
		TursoDbUrl:       os.Getenv("TURSO_DATABASE_URL"),
		TursoDBAuthToken: os.Getenv("TURSO_AUTH_TOKEN"),
	}
	return configuration
}
