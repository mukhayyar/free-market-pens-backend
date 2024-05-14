package user_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	userIdStr := c.Param("userId")

	if userIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "userId is required"})
	}
	
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}
	
	result, err := models.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateUser(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	whatsappNumber := c.FormValue("whatsappNumber")
	fullName := c.FormValue("fullName")
	password := c.FormValue("password")

	result, err := models.CreateUser(email, username, whatsappNumber, fullName, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
