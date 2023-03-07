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
	g.POST("/signup", signup)
	g.POST("/signin", signin)
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

func signin(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	id, err := user.Signin()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Password or id is wrong!")
	}

	fmt.Println("authenticated user id: ", id)
	return c.String(http.StatusOK, "Successfully signed in!")
}
