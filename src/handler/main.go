// Package handler is bussiness logic for server
package handler

import (
	"echo-todo-server/src/model"
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
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	if len(todos) == 0 {
		return c.String(http.StatusNotFound, "Todo was not found!")
	}

	return c.JSON(http.StatusOK, todos)
}

// CreateTodo is for making a new Todo
func CreateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	if err := todo.CreateTodo(); err != nil {
		return c.String(http.StatusInternalServerError, "Invalid value.")
	}

	return c.String(http.StatusOK, "Successfully added todo!")
}

// UpdateTodo is for update todo
func UpdateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	id := c.Param("id")
	if err := todo.UpdateTodo(id); err != nil {
		return c.String(http.StatusBadRequest, "Invalid id.")
	}

	return c.String(http.StatusOK, "Update todo was succeeded!")
}

// DeleteTodo is for delete todo
func DeleteTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	id := c.Param("id")
	if err := todo.DeleteTodo(id); err != nil {
		return c.String(http.StatusBadRequest, "Invalid id.")
	}

	return c.String(http.StatusOK, "Successfully deleted todo!")
}
