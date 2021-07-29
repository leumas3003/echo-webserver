package api

import (
	"github.com/labstack/echo/v4"
	"github.com/leumas3003/echo-webserver/api/handlers"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookies)
}
