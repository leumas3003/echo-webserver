package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You dont have any cookie")
			}
			log.Println(err)
			return err
		}
		//Fake check, we must check with some DB
		if cookie.Value == "some_string" {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "You dont have the right cookie")
	}
}
