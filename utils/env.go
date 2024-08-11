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
		log.Print(err)
	}

	e.GoEnv = populateEnvVar("GO_ENV")
	e.LoggerUrl = populateEnvVar("LOGGER_URL")
	e.DatabaseURL = populateEnvVar("DATABASE_URL")
	e.DatabaseUser = populateEnvVar("DATABASE_USER")
	e.DatabasePassword = populateEnvVar("DATABASE_PASSWORD")

	if e.LoggerUrl == "" {
		log.Panic("Error loading LOGGER_URL env variable")
	}
}

func populateEnvVar(varName string) string {
	result := os.Getenv(varName)
	if result == "" {
		log.Panic("Error loading LOGGER_URL env variable")
	}
	return result
}

func (e Environment) IsDevMode() bool {
	return e.GoEnv == "development"
}
