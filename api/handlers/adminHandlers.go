package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MainAdmin godoc
// @Summary MainAdmin
// @Description Access to mainadmin page
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /admin/main [get]
func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret page")
}
