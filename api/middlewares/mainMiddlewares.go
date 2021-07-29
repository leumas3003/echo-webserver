package middlewares

import (
	"github.com/labstack/echo/v4"
)

func SetMainMiddlewares(e *echo.Echo) {
	//e.Use(middleware.Static("./static"))
	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("JUSTATEST", "ITSWORK")
		return next(c)
	}
}
