package group

import (
	"echo-todo-server/src/model"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group) {
	g.POST("signup", signup)
	// g.GET("", get)
	// g.PUT(":id", update)
	// g.DELETE(":id", delete)
}

var validate *validator.Validate

func signup(c echo.Context) error {
	user := new(model.User)
	// var user model.User
	if err := c.Bind(user); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	fmt.Println(user)
	if err := user.Signup(); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Invalid value.")
	}

	return c.String(http.StatusOK, "Successfully created user")
}
