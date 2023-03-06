// Package env is to check for missing definitions.
package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DBNAME   string
)

func getEnv(envName string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		log.Printf("warn: %s is not valid\n", envName)
	}
	return envValue
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_HOST = getEnv("POSTGRES_HOST")
	POSTGRES_USER = getEnv("POSTGRES_USER")
	POSTGRES_PORT = getEnv("POSTGRES_PORT")
	POSTGRES_PASSWORD = getEnv("POSTGRES_PASSWORD")
	POSTGRES_DBNAME = getEnv("POSTGRES_DBNAME")
}
