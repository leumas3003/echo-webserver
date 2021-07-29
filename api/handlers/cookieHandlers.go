package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MainCookie godoc
// @Summary MainCookie
// @Description Access to MainCookie page
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /cookie/main [get]
func MainCookies(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret page")
}
