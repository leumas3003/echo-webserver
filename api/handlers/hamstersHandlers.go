package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// AddHamster godoc
// @Summary AddHamster
// @Description AddHamster
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Param hamster body Hamster true "hamster"
// @Router /hamsters [post]
func AddHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing AddHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}
