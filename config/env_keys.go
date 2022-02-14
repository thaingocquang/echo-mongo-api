package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Database struct {
	URI  string
	Name string
}

// ENV .env struct
type ENV struct {
	// App port
	AppPort string

	// Database
	Database Database
}

var env ENV

// InitDotEnv init params in .env file
func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	appPort := os.Getenv("APP_PORT")
	database := Database{URI: os.Getenv("DB_URI"), Name: os.Getenv("DB_Name")}
	env = ENV{
		AppPort:  appPort,
		Database: database,
	}
}

// GetEnv return .env data
func GetEnv() *ENV {
	return &env
}
