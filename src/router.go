// Package router is routing
package router

import (
	"echo-todo-server/src/group"

	"github.com/labstack/echo/v4"
)

// New is a router
func New() *echo.Echo {
	e := echo.New()

	group.Todo(e.Group("/todo"))

	return e
}
