package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	//this logs the server interaction
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + ";\n",
	}))

	//Basic Auth
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		//check in the DB
		if username == "user" && password == "pass" {
			return true, nil
		}
		msg := echo.Map{"message": "Unauthorized"}
		err := c.JSON(http.StatusUnauthorized, msg)
		return false, err
	}))
}
