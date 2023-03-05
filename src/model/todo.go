package model

import (
	"log"
	"time"
)

// CreateTodoTable is a function for creating todo table
func CreateTodoTable() (err error) {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Todo is todo struct
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" varidate:"required, min=1, max=255"`
	Description string    `json:"description" varidate:"required  min=1 max=255"`
	CreatedAt   time.Time `json:"created_at"`
}

// Read is a fetching todos
func Read() (todos []Todo, err error) {
	query := `SELECT id, title, description, created_at FROM todos;`

	rows, err := Db.Query(query)
	if err != nil {
		log.Println("Error Get todo list: ", err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// Create is used to add a new todo to the database.
func (t *Todo) Create() (err error) {
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

// Update is used to update a todo
func (t *Todo) Update(id string) (err error) {
	query := `UPDATE todos SET title = $1, description = $2 WHERE id = $3`

	if _, err = Db.Exec(query,
		t.Title,
		t.Description,
		id,
	); err != nil {
		log.Println("Error UPDATE todo: ", err)
	}
	return err
}

// Delete is used to delete a todo
func (t *Todo) Delete(id string) error {
	query := `DELETE FROM todos WHERE id = $1`

	if _, err := Db.Exec(query,
		id,
	); err != nil {
		log.Println("Error DELETE todo: ", err)
	}

	return err
}
