// Package model is for database
package model

import (
	"database/sql"
	"echo-todo-server/src/env"
	"fmt"
	"log"

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
		env.POSTGRES_HOST,
		env.POSTGRES_PORT,
		env.POSTGRES_USER,
		env.POSTGRES_PASSWORD,
		env.POSTGRES_DBNAME,
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

	// create table
	CreateTodoTable()
	CreateUserTable()
	CreateSessionTable()
}
