// Package handler is bussiness logic for server
package handler

import (
	"echo-todo-server/src/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseHelloWorld is for testing for responses of api
func ResponseHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

// GetTodoList returns a list todo for a repository
func GetTodoList(c echo.Context) error {
	todos, err := model.SelectTodoList()
	if err != nil {
		fmt.Printf("Failed to reading the todo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, todos)
}

// CreateTodo is for making a new Todo
func CreateTodo(c echo.Context) error {
	fmt.Println("create todo")
	todo := model.Todo{}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to reading the request body for addTodo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if err = json.Unmarshal(body, &todo); err != nil {
		log.Printf("Failed unmarshaling n CreateTodo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if err = todo.CreateTodo(); err != nil {
		log.Printf("Failed creating todo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	fmt.Printf("this is your todo: %#v", todo)
	return c.String(http.StatusOK, "Create todo was succeeded!")
}

// UpdateTodo is for update todo
func UpdateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		log.Printf("Failed binding to todo for update: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if err := todo.UpdateTodo(); err != nil {
		log.Printf("Failed updateing todo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, "Update todo was succeeded!")
}

// DeleteTodo is for delete todo
func DeleteTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		log.Printf("Failed binding to todo for delete: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	if err := todo.DeleteTodo(); err != nil {
		log.Printf("Failed deleting todo: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, "Delete todo was succeeded!")
}
