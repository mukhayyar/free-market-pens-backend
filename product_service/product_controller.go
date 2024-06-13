package product_controller

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

// Claims represents the JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

const imgbbAPIKey = "a818e1c105d4ad0fa46f04cf1e30c957"

type imgbbResponse struct {
	Data struct {
		DisplayURL string `json:"display_url"`
	} `json:"data"`
}

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

	if productIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "productId is required"})
	}

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}

	productDetailResult, err := models.GetMyProductDetail(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	batchResult, err := models.GetLastBatch(productId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	fmt.Println(productDetailResult.Data.(map[string]interface{}))
	// Assuming productDetailResult contains a storeId field
	productData := productDetailResult.Data.(map[string]interface{})["product"].(map[string]interface{})
	storeId := productData["storeId"].(int)

	storeResult, err := models.GetMyStore(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := map[string]interface{}{
		"productDetailResult": productDetailResult,
		"batchResult":         batchResult.Data,
		"storeResult":         storeResult,
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllMyProduct(c echo.Context) error {
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

	result, err := models.GetAllMyProduct(storeId)
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

func CreateProduct(c echo.Context) error {
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

	name := c.FormValue("name")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	if name == "" || description == "" || file == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "data can't be empty"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to open image file"})
	}
	defer src.Close()

	photoURL, err := uploadImageToImgbb(src, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload image"})
	}

	result, err := models.CreateProduct(storeId, photoURL, name, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
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

	productIdStr := c.Param("productId")
	name := c.FormValue("name")
	description := c.FormValue("description")
	file, err := c.FormFile("photo")

	if productIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "productId is required"})
	}

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid productId"})
	}

	var photoURL string
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to open image file"})
		}
		defer src.Close()

		photoURL, err = uploadImageToImgbb(src, file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to upload image"})
		}
	}

	result, err := models.UpdateProduct(productId, photoURL, name, description)
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
