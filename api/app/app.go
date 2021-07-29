package app

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/leumas3003/echo-webserver/docs"
	"github.com/leumas3003/echo-webserver/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email samuel.martel@derivco.se

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func StartServer() *echo.Echo {
	fmt.Println("Starting echo server")
	e := router.New()
	//Starting the webserver
	e.Start(":3001")
	confSwagger()
	return e
}

func confSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Animals API"
	docs.SwaggerInfo.Description = "Webserver using echo. Basic Auth, JWT, Log,."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
