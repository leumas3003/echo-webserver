package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// AddCat godoc
// @Summary AddCat
// @Description AddCat
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Param cat body Cat true "cat"
// @Router /cats [post]
func AddCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for AddCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshal in AddCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got your cat!")
}

// GetCats godoc
// @Summary GetCats
// @Description GetCats
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Param name query string true "name"
// @Param type query string true "type"
// @Param data path string true "data"
// @Router /cats/{data} [get]
func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("The cat name is %s and the type %s ", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to choose string or json",
	})
}
