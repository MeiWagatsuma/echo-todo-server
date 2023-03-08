package group

import (
	"echo-todo-server/src/model"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Auth(g *echo.Group) {
	g.POST("/signup", signup)
	g.POST("/signin", signin)
	g.DELETE("/signout", signout)
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

	userId, err := user.Signin()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Password or name is wrong!")
	}
	session := model.Session{UserId: userId}

	sessionExists, err := session.Exists(userId)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	if sessionExists {
		return c.String(http.StatusConflict, "Already signed in")
	}

	if err := session.Generate(); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    session.Token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusOK, "Successfully signed in")
}

func signout(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		if errors.Is(err, http.ErrNoCookie) {
			return c.JSON(http.StatusOK, "You are already signouted")
		}

		return c.String(http.StatusInternalServerError, "Invalid value")
	}

	session := model.Session{Token: cookie.Value}

	if err := session.Delete(); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Oops! Something went wrong!")
	}

	// Delete client cookie
	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusOK, "Successfully signouted in")
}
