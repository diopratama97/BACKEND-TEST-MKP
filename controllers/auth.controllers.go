package controllers

import (
	"backend-test-mkp/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.Login(email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"email": email,
		"token": t,
	})
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	telp := c.FormValue("telp")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")

	if role != "ADMIN" && role != "USER" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Role must be ADMIN or USER"})
	}

	result, err := models.Register(email, password, address, telp, role, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
