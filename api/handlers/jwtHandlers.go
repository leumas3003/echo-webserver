package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MainJWT godoc
// @Summary MainJWT
// @Description Access to MainJWT page
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /jwt/main [get]
func MainJwt(c echo.Context) error {
	/*Check this topic because the user.(*jwt.Token) line fails
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	log.Println("User Name:", claims["name"], "User ID: ", claims["jti"])
	*/
	return c.String(http.StatusOK, "you are on the top secret jwt page")
}
