package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func middlewareAdapter(e *echo.Echo) {
	e.Use(middleware.BodyLimit("2K"))
}
