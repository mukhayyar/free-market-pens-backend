package product_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllMyProduct(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}

	result, err := models.GetAllMyProduct(storeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateProduct(c echo.Context) error {
	storeId := c.FormValue("storeId")
	photo := c.FormValue("photo")
	name := c.FormValue("name")
	description := c.FormValue("description")

	if storeId == "" && photo == "" && name == "" && description == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}
	
	result, err := models.CreateProduct(storeId, photo, name, description)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}