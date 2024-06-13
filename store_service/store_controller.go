package store_controller

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const imgbbAPIKey = "a818e1c105d4ad0fa46f04cf1e30c957"

type imgbbResponse struct {
	Data struct {
		DisplayURL string `json:"display_url"`
	} `json:"data"`
}

// Claims represents the JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func CheckUserHasStore(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "userId is required",
			"data":    nil,
		})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid userId",
			"data":    nil,
		})
	}

	hasStore, err := models.CheckUserHasStore(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": "Error checking user store",
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "User store check successful",
		"data":    hasStore,
	})
}

func GetMyStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "storeId is required",
			"data":    nil,
		})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid storeId",
			"data":    nil,
		})
	}

	result, err := models.GetMyStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "Store retrieved successfully",
		"data":    result,
	})
}

func GetStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "storeId is required",
			"data":    nil,
		})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid storeId",
			"data":    nil,
		})
	}

	result, err := models.GetStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "Store retrieved successfully",
		"data":    result,
	})
}

func uploadImageToImgbb(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", fileHeader.Filename)
	if err != nil {
		return "", err
	}
	part.Write(fileBytes)
	writer.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.imgbb.com/1/upload?key=%s", imgbbAPIKey), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var imgbbResp imgbbResponse
	if err := json.Unmarshal(respBody, &imgbbResp); err != nil {
		return "", err
	}

	return imgbbResp.Data.DisplayURL, nil
}

func CreateStore(c echo.Context) error {
	userIdStr := c.FormValue("userId")
	name := c.FormValue("name")
	whatsappNumber := c.FormValue("whatsappNumber")
	file, err := c.FormFile("photoProfile")

	if userIdStr == "" || name == "" || whatsappNumber == "" || file == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "userId, name, whatsappNumber, and photoProfile are required",
			"data":    nil,
		})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid userId",
			"data":    nil,
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": "Unable to open image file",
			"data":    nil,
		})
	}
	defer src.Close()

	photoProfileURL, err := uploadImageToImgbb(src, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": "Failed to upload image",
			"data":    nil,
		})
	}

	result, err := models.CreateStore(userId, name, photoProfileURL, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "Store created successfully",
		"data":    result,
	})
}

func UpdateStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	name := c.FormValue("name")
	whatsappNumber := c.FormValue("whatsappNumber")
	file, err := c.FormFile("photoProfile")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "storeId is required",
			"data":    nil,
		})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid storeId",
			"data":    nil,
		})
	}

	var photoProfileURL string
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"status":  http.StatusInternalServerError,
				"message": "Unable to open image file",
				"data":    nil,
			})
		}
		defer src.Close()

		photoProfileURL, err = uploadImageToImgbb(src, file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"status":  http.StatusInternalServerError,
				"message": "Failed to upload image",
				"data":    nil,
			})
		}
	}

	result, err := models.UpdateStore(storeId, name, photoProfileURL, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "Store updated successfully",
		"data":    result,
	})
}

func CloseStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "storeId is required",
			"data":    nil,
		})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"status":  http.StatusBadRequest,
			"message": "Invalid storeId",
			"data":    nil,
		})
	}

	result, err := models.CloseStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"status":  http.StatusOK,
		"message": "Store closed successfully",
		"data":    result,
	})
}
