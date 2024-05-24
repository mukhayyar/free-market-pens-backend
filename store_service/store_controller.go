package store_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMyStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}
	
	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.GetMyStore(storeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}
	
	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.GetStore(storeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateStore(c echo.Context) error {
	userIdStr := c.FormValue("userId")
	name := c.FormValue("name")
	photoProfile := c.FormValue("photoProfile")
	whatsappNumber := c.FormValue("whatsappNumber")

	if userIdStr == "" || name == "" || photoProfile == "" || whatsappNumber == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	result, err := models.CreateStore(userId, name, photoProfile, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	name := c.FormValue("name")
	photoProfile := c.FormValue("photoProfile")
	whatsappNumber := c.FormValue("whatsappNumber")
	
	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}
	
	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.UpdateStore(storeId, name, photoProfile, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

func CloseStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}

	result, err := models.CloseStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}