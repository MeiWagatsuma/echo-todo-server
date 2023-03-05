// Package router is routing
package router

import (
	"echo-todo-server/src/group"
	"echo-todo-server/src/handler"

	"github.com/labstack/echo/v4"
)

// New is a router
func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.ResponseHelloWorld)
	group.Todo(e.Group("/todo"))

	return e
}
