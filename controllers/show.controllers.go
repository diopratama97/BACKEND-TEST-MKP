package controllers

import (
	"backend-test-mkp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListShow(c echo.Context) error {
	result, err := models.ListShow()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateShow(c echo.Context) error {
	name := c.FormValue("name")
	bioskop_id := c.FormValue("bioskop_id")
	show_date := c.FormValue("show_date")
	show_time := c.FormValue("show_time")
	price := c.FormValue("price")
	status := c.FormValue("status")
	url_image := c.FormValue("url_image")
	url_trailer := c.FormValue("url_trailer")

	if status != "READY" && status != "COMINGSOON" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Status only READY or COMINGSOON"})
	}

	newPrice, _ := strconv.Atoi(price)

	result, err := models.CreateShow(newPrice, bioskop_id, name, show_date, show_time, status, url_trailer, url_image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// func UpdateUsers(c echo.Context) error {
// 	id := c.Param("id")
// 	// name := c.FormValue("name")
// 	// address := c.FormValue("address")
// 	// telp := c.FormValue("telp")

// 	conv_id, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	result, err := models.UpdateUsers(conv_id, name, address, telp)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)

// }

// func DeleteUsers(c echo.Context) error {
// 	id := c.Param("id")

// 	conv_id, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	result, err := models.DeleteUsers(conv_id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)

// }

// func DetailUsers(c echo.Context) error {
// 	id := c.Param("id")

// 	conv_id, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	result, err := models.DetailUsers(conv_id)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, result)

// }
