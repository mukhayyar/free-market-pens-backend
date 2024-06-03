package store_pickup_place_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllStorePickupPlace(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}

	result, err := models.GetAllStorePickupPlace(storeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateStorePickupPlace(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	name := c.FormValue("name")

	if storeIdStr == "" || name == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.CreateStorePickupPlace(storeId, name)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

func UpdateStorePickupPlace(c echo.Context) error {
	storePickupPlaceIdStr := c.Param("storePickupPlaceId")
	storeIdStr := c.Param("storeId")
	name := c.FormValue("name")

	if storeIdStr == "" || storePickupPlaceIdStr == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	storePickupPlaceId, err := strconv.Atoi(storePickupPlaceIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storePickupPlaceId"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.UpdateStorePickupPlace(storePickupPlaceId, storeId, name)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

func DeleteStorePickupPlace(c echo.Context) error {
	storePickupPlaceIdStr := c.Param("storePickupPlaceId")

	if storePickupPlaceIdStr == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "storePickupPlaceId can't be empty"})
	}

	storePickupPlaceId, err := strconv.Atoi(storePickupPlaceIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storePickupPlaceId"})
	}

	
	result, err := models.DeleteStorePickupPlace(storePickupPlaceId)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}