// Package router is routing
package router

import (
	"echo-todo-server/src/group"
	"net/http"

	"github.com/labstack/echo/v4"
)

// New is a router
func New() *echo.Echo {
	e := echo.New()

	middlewareAdapter(e)

	group.Todo(e.Group("/todo", cookieTokenAuthMiddleware))
	group.Auth(e.Group("/auth"))

	return e
}

const authToken = "pass"

func cookieTokenAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error { // Cookieからトークンを取得する
		cookie, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "トークンが提供されていません")
		}

		// トークンが正しい場合は、次のハンドラーを呼び出す
		if cookie.Value == authToken {
			return next(c)
		}

		// トークンが正しくない場合は、エラーを返す
		return echo.NewHTTPError(http.StatusUnauthorized, "無効なトークンが提供されました")
	}
}
