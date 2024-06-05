package user_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Claims represents the JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

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
	password := c.FormValue("password")

	if email == "" || username == "" || password == "" || whatsappNumber == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	result, err := models.CreateUser(email, username, whatsappNumber, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
