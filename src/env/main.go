// Package env is to check for missing definitions.
package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DBNAME   string

	SESSION_KEY_LENGTH int
)

func getEnv(envName string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		log.Printf("warn: %s is not valid\n", envName)
	}
	return envValue
}

var err error

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_HOST = getEnv("POSTGRES_HOST")
	POSTGRES_USER = getEnv("POSTGRES_USER")
	POSTGRES_PORT = getEnv("POSTGRES_PORT")
	POSTGRES_PASSWORD = getEnv("POSTGRES_PASSWORD")
	POSTGRES_DBNAME = getEnv("POSTGRES_DBNAME")

	if SESSION_KEY_LENGTH, err = strconv.Atoi(getEnv("SESSION_KEY_LENGTH")); err != nil {
		log.Fatal("SESSION_KEY_LENGTH must be a number")
	}
}
