// Package handler is bussiness logic for server
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseHelloWorld is for testing for responses of api
func ResponseHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
