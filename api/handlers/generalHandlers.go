package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hi godoc
// @Summary HI
// @Description Access to Hi page
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /hi [get]
func Hi(c echo.Context) error {
	return c.String(http.StatusOK, "hello from the other side")
}
