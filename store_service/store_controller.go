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
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "data can't be empty"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to open image file"})
	}
	defer src.Close()

	photoProfileURL, err := uploadImageToImgbb(src, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload image"})
	}

	result, err := models.CreateStore(userId, name, photoProfileURL, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStore(c echo.Context) error {
	storeIdStr := c.Param("storeId")
	name := c.FormValue("name")
	whatsappNumber := c.FormValue("whatsappNumber")
	file, err := c.FormFile("photoProfile")

	if storeIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "storeId is required"})
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid storeId"})
	}

	var photoProfileURL string
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to open image file"})
		}
		defer src.Close()

		photoProfileURL, err = uploadImageToImgbb(src, file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload image"})
		}
	}

	result, err := models.UpdateStore(storeId, name, photoProfileURL, whatsappNumber)
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
