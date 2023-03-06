// Package group is usecases
package group

import (
	"echo-todo-server/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Todo has all todo methods
func Todo(g *echo.Group) {
	g.GET("", getTodoList)
	g.POST("", createTodo)
	g.PUT(":id", updateTodo)
	g.DELETE(":id", deleteTodo)
}

// GetTodoList returns a list todo for a repository
func getTodoList(c echo.Context) error {
	todos, err := model.Read()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	if len(todos) == 0 {
		return c.String(http.StatusNotFound, "Todo was not found!")
	}

	return c.JSON(http.StatusOK, todos)
}

// CreateTodo is for making a new Todo
func createTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	if err := todo.Create(); err != nil {
		return c.String(http.StatusInternalServerError, "Invalid value.")
	}

	return c.String(http.StatusOK, "Successfully added todo!")
}

// UpdateTodo is for update todo
func updateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	id := c.Param("id")
	if err := todo.Update(id); err != nil {
		return c.String(http.StatusBadRequest, "Invalid id.")
	}

	return c.String(http.StatusOK, "Update todo was succeeded!")
}

// DeleteTodo is for delete todo
func deleteTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	id := c.Param("id")
	if err := todo.Delete(id); err != nil {
		return c.String(http.StatusBadRequest, "Invalid id.")
	}

	return c.String(http.StatusOK, "Successfully deleted todo!")
}
