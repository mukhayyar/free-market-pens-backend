package cart_controller

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

func GetCartByID(c echo.Context) error {
	cartIDStr := c.Param("cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid cart ID"})
	}

	res, err := models.GetCart(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func GetCartsByUserID(c echo.Context) error {
	// Extract user ID from JWT token claims
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(Claims)
	userID := claims.UserID

	res, err := models.GetCartsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func CreateCart(c echo.Context) error {
	userIDStr := c.FormValue("user_id")
	productIDStr := c.FormValue("product_id")
	quantityStr := c.FormValue("quantity")

	if userIDStr == "" || productIDStr == "" || quantityStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "data can't be empty"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid product ID"})
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid quantity"})
	}

	res, err := models.CreateCart(userID, productID, quantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateCart(c echo.Context) error {
	cartIDStr := c.Param("cartID")
	quantityStr := c.FormValue("quantity")

	if cartIDStr == "" || quantityStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "data can't be empty"})
	}

	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid cart ID"})
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid quantity"})
	}

	res, err := models.UpdateCart(cartID, quantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteCart(c echo.Context) error {
	cartIDStr := c.Param("cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid cart ID"})
	}

	res, err := models.DeleteCart(cartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
