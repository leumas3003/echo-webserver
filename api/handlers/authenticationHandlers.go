package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/leumas3003/echo-webserver/api/utils/errors"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type LoginStruct struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Login godoc
// @Summary Login
// @Description Login by basic auth
// @Accept json
// @Produce json
// @Success 200 {object} LoginStruct
// @Failure 401 {object} errors.ApiError
// @Failure 500 {object} errors.ApiError
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Router /login [get]
func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	//check username and password against DB after hashing the password
	if username == "user" && password == "pass" {
		//this is the same
		//cookie := new(http.Cookie)
		cookie := &http.Cookie{
			Name:    "sessionID",
			Value:   "some_string",
			Expires: time.Now().Add(48 * time.Hour),
		}
		c.SetCookie(cookie)

		//TODO: create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT Token", err)
			return c.JSON(http.StatusInternalServerError, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: "something went wrong",
				Error:   err.Error(),
			})
		}

		jwtCookie := &http.Cookie{
			Name:    "JWTCookie",
			Value:   token,
			Expires: time.Now().Add(48 * time.Hour),
		}
		c.SetCookie(jwtCookie)
		return c.JSON(http.StatusOK, &LoginStruct{
			Message: "You were logged in!",
			Token:   token,
		})
	}
	return c.JSON(http.StatusInternalServerError, &errors.ApiError{
		Status:  http.StatusUnauthorized,
		Message: "Your username or password were wrong",
	})
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"samuel",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	//NEVER EVER DO THIS. YOU MUST HAVE A LONG SECRET KEY
	token, err := rawToken.SignedString([]byte("mySecretKey"))
	if err != nil {
		return "", err
	}
	return token, nil
}
