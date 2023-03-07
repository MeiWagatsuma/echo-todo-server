package model

import (
	"echo-todo-server/src/lib"
	"log"
)

func CreateUserTable() (err error) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// User is user struct
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" varidate:"required"`
	Password string `json:"password" varidate:"required"`
}

func (u *User) Signup() (err error) {
	query := "INSERT INTO users (name, password) VALUES($1, $2)"

	hashedPassword := lib.Sha3Hash(u.Password)
	_, err = Db.Exec(query, u.Name, hashedPassword)
	if err != nil {
		log.Println("Error create user: ", err)
	}

	return err
}
