package router

import (
	"github.com/labstack/echo/v4"
	"github.com/leumas3003/echo-webserver/api"
	"github.com/leumas3003/echo-webserver/api/middlewares"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() *echo.Echo {
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")
	cookiesGroup := e.Group("/cookies")
	jwtGroup := e.Group("/jwt")
	swagger := e.Group("/swagger")

	//Swagger
	swagger.GET("/*", echoSwagger.WrapHandler)
	//set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookiesGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	//set main routes
	api.MainGroup(e)
	//set admin group
	api.AdminGroup(adminGroup)
	//set cookie group
	api.CookieGroup(cookiesGroup)
	//set jwt group
	api.JwtGroup(jwtGroup)

	return e
}
