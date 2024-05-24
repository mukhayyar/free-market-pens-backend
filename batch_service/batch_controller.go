package batch_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBatch(c echo.Context) error {
	productIdStr := c.Param("productId")
	pickupPlaceIdStr := c.FormValue("pickupPlaceId")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	pickupDate := c.FormValue("pickupDate")
	pickupTime := c.FormValue("pickupTime")
	closeOrderDate := c.FormValue("closeOrderDate")
	closeOrderTime := c.FormValue("closeOrderTime")

	if productIdStr == "" || pickupPlaceIdStr == "" || stockStr == "" || priceStr == "" || pickupDate == "" || pickupTime == "" || closeOrderDate == "" || closeOrderTime == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}

	pickupPlaceId, err := strconv.Atoi(pickupPlaceIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid pickupPlaceId"})
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid stock"})
	}

	price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid price"})
    }

	pickupTime = pickupDate + " " + pickupTime
    closeOrderTime = closeOrderDate + " " + closeOrderTime

	batchResult, err := models.CreateBatch(productId, pickupPlaceId, stock, price, pickupTime, closeOrderTime)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := map[string]interface{}{
		"batchResult": batchResult,
	}
	
	return c.JSON(http.StatusOK, response)
}

func UpdateBatch(c echo.Context) error {
	batchIdStr := c.Param("batchId")
	pickupPlaceIdStr := c.FormValue("pickupPlaceId")
	stockStr := c.FormValue("stock")
	priceStr := c.FormValue("price")
	pickupTime := c.FormValue("pickupTime")
	closeOrderTime := c.FormValue("closeOrderTime")
	var pickupPlaceId int
	var stock int
	var price float64
	
	if batchIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "batchId is required"})
	}
	
	batchId, err := strconv.Atoi(batchIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid batchId"})
	}

	if pickupPlaceIdStr == "" {
        pickupPlaceId = 0
    } else {
        pickupPlaceId, err = strconv.Atoi(pickupPlaceIdStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid pickupPlaceId"})
        }
    }

    if stockStr == "" {
        stock = 0
    } else {
        stock, err = strconv.Atoi(stockStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid stock"})
        }
    }

    if priceStr == "" {
        price = 0
    } else {
        price, err = strconv.ParseFloat(priceStr, 64)
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid price"})
        }
    }
	
	result, err := models.UpdateBatch(batchId, pickupPlaceId, stock, price, pickupTime, closeOrderTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK, result)
}