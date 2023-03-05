// Package group is usecases
package group

import (
	"echo-todo-server/src/handler"

	"github.com/labstack/echo/v4"
)

// Todo has all todo methods
func Todo(g *echo.Group) {
	g.GET("", handler.GetTodoList)
	g.POST("", handler.CreateTodo)
	g.PUT(":id", handler.UpdateTodo)
	g.DELETE(":id", handler.DeleteTodo)
}
