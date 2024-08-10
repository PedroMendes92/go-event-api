package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Env Environment

type Environment struct {
	GoEnv            string
	LoggerUrl        string
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
}

func (e *Environment) InitEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.GoEnv = os.Getenv("GO_ENV")
	e.LoggerUrl = os.Getenv("LOGGER_URL")
	e.DatabaseURL = os.Getenv("DATABASE_URL")
	e.DatabaseUser = os.Getenv("DATABASE_USER")
	e.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
}

func (e Environment) IsDevMode() bool {
	return e.GoEnv == "development"
}
