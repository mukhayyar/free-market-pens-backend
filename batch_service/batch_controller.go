package batch_controller

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBatch(c echo.Context) error {
	productIdStr := c.FormValue("productId")
	pickupPlaceIdStr := c.FormValue("pickupPlaceId")
	priceStr := c.FormValue("price")
	stockStr := c.FormValue("stock")
	pickupDate := c.FormValue("pickupDate")
	pickupTime := c.FormValue("pickupTime")
	closeOrderDate := c.FormValue("closeOrderDate")
	closeOrderTime := c.FormValue("closeOrderTime")

	if productIdStr == "" && pickupPlaceIdStr == "" && stockStr == "" && pickupDate == "" && pickupTime == "" && closeOrderDate == "" && closeOrderTime == "" {
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
	
	priceResult, err := models.AddPrice(productId, price)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	batchResult, err := models.CreateBatch(productId, pickupPlaceId, stock, pickupDate, pickupTime, closeOrderDate, closeOrderDate)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := map[string]interface{}{
		"priceResult": priceResult,
		"batchResult": batchResult,
	}
	
	return c.JSON(http.StatusOK, response)
}