package store_pickup_place_controller

import (
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GetAllStorePickupPlace(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing token"})
	}

	tokenStr := authHeader[len("Bearer "):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}

	userId := claims.UserID

	// Fetch the storeId associated with the userId
	storeResult, err := models.GetStoreIdByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get store ID"})
	}

	storeId, ok := storeResult.Data.(map[string]int)["store_id"]
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse store ID"})
	}

	result, err := models.GetAllStorePickupPlace(storeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateStorePickupPlace(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing token"})
	}

	tokenStr := authHeader[len("Bearer "):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}

	userId := claims.UserID

	// Fetch the storeId associated with the userId
	storeResult, err := models.GetStoreIdByUserId(userId)
	fmt.Println(storeResult)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get store ID"})
	}

	storeId, ok := storeResult.Data.(map[string]int)["store_id"]
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse store ID"})
	}

	name := c.FormValue("name")
	if name == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	result, err := models.CreateStorePickupPlace(storeId, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStorePickupPlace(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing token"})
	}

	tokenStr := authHeader[len("Bearer "):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}

	userId := claims.UserID

	// Fetch the storeId associated with the userId
	storeResult, err := models.GetStoreIdByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get store ID"})
	}

	storeId, ok := storeResult.Data.(map[string]int)["store_id"]
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse store ID"})
	}

	storePickupPlaceIdStr := c.Param("storePickupPlaceId")
	name := c.FormValue("name")

	if storePickupPlaceIdStr == "" || name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storePickupPlaceId and name are required"})
	}

	storePickupPlaceId, err := strconv.Atoi(storePickupPlaceIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storePickupPlaceId"})
	}

	result, err := models.UpdateStorePickupPlace(storePickupPlaceId, storeId, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteStorePickupPlace(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing token"})
	}

	tokenStr := authHeader[len("Bearer "):]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}
	if !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}

	userId := claims.UserID

	// Fetch the storeId associated with the userId
	storeResult, err := models.GetStoreIdByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get store ID"})
	}

	storeId, ok := storeResult.Data.(map[string]int)["store_id"]
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse store ID"})
	}

	fmt.Println(storeId)

	storePickupPlaceIdStr := c.Param("storePickupPlaceId")
	if storePickupPlaceIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storePickupPlaceId can't be empty"})
	}

	storePickupPlaceId, err := strconv.Atoi(storePickupPlaceIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storePickupPlaceId"})
	}

	result, err := models.DeleteStorePickupPlace(storePickupPlaceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
