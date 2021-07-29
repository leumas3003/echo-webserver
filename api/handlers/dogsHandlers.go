package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// AddDog godoc
// @Summary AddDog
// @Description AddDog
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Param dog body Dog true "dog"
// @Router /dogs [post]
func AddDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing AddDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")

}
