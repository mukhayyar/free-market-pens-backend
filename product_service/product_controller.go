package product_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllProduct(c echo.Context) error {
	closeOrderDate := c.FormValue("closeOrderDate")
	pickupDate := c.FormValue("pickupDate")

	result, err := models.GetAllProduct(closeOrderDate, pickupDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMyProductDetail(c echo.Context) error {
	productIdStr := c.Param("productId")

	if productIdStr == ""{
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "productId is required"})
	}

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}
	
	productDetailResult, err := models.GetMyProductDetail(productId)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	batchResult, err := models.GetAllBatch(productId)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := map[string]interface{}{
		"productDetailResult": productDetailResult,
		"batchResult": batchResult,
	}
	
	return c.JSON(http.StatusOK, response)
}

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
	storeIdStr := c.Param("storeId")
	photo := c.FormValue("photo")
	name := c.FormValue("name")
	description := c.FormValue("description")

	if storeIdStr == "" || photo == "" || name == "" || description == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}
	
	result, err := models.CreateProduct(storeId, photo, name, description)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	productIdStr := c.Param("productId")
	photo := c.FormValue("photo")
	name := c.FormValue("name")
	description := c.FormValue("description")
	
	if productIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "productId is required"})
	}
	
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}
	
	result, err := models.UpdateProduct(productId, photo, name, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	productIdStr := c.Param("productId")

	if productIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "productId is required"})
	}

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}

	result, err := models.DeleteProduct(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}