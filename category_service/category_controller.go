package category_controller

import (
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	result, err := models.GetAllCategory()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

func CreateCategory(c echo.Context) error {
	name := c.FormValue("name")

	if name == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}
	
	result, err := models.CreateCategory(name)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}