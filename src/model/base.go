// Package model is for database
package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres driver
)

// Db is database struct
var Db *sql.DB

var err error

func init() {
	err = godotenv.Load(".env")

	cmd := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	Db, err = sql.Open("postgres", cmd)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database.")

	defer Db.Close()
	defer fmt.Println("close")
}
