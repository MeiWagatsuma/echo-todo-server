// Package router is routing
package router

import (
	"echo-todo-server/src/handler"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.ResponseHelloWorld)

	return e
}
