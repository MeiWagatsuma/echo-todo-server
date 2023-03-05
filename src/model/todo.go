package model

import (
	"log"
	"time"
)

func CreateTodoTable() (err error) {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

type Todo struct {
	Title       string    `json:title`
	Description string    `json:description`
	CreatedAt   time.Time `created_at`
}

func (t *Todo) CreateTodo() (err error) {
	cmd := `INSERT INTO todos (title, description, created_at) VALUES($1, $2, $3 )`

	_, err = Db.Exec(cmd,
		t.Title,
		t.Description,
		time.Now(),
	)

	if err != nil {
		log.Println("Error insert todo: ", err)
	}

	return err
}
