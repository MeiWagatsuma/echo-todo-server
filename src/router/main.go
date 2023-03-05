// Package router is routing
package router

import (
	"echo-todo-server/src/handler"

	"github.com/labstack/echo/v4"
)

// New is a router
func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.ResponseHelloWorld)
	e.GET("/todos", handler.GetTodoList)
	e.POST("/todos", handler.CreateTodo)
	e.PUT("/todos", handler.UpdateTodo)
	e.DELETE("/todos", handler.DeleteTodo)

	return e
}
