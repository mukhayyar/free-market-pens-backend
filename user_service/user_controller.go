package user_controller

import (
	"backend/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type User struct {
	UserId         int
	Email          string
	Username       string
	WhatsappNumber string
	Password       string
	FullName       string
	CreatedAt      string
	UpdatedAt      string
	IsAdmin        bool
}

// Claims represents the JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GetProfile(c echo.Context) error {
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

	result, err := models.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetUser(c echo.Context) error {
	userIdStr := c.Param("userId")

	if userIdStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "userId is required"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	result, err := models.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateUser(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	whatsappNumber := c.FormValue("whatsappNumber")
	password := c.FormValue("password")

	if email == "" || username == "" || password == "" || whatsappNumber == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "data can't be empty"})
	}

	result, err := models.CreateUser(email, username, whatsappNumber, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllUsers(c echo.Context) error {
	result, err := models.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
func UpdatePassword(c echo.Context) error {
	userIdStr := c.FormValue("userId")
	// oldPassword := c.FormValue("oldPassword")
	newPassword := c.FormValue("newPassword")
	confirmNewPassword := c.FormValue("confirmNewPassword")

	if userIdStr == "" || newPassword == "" || confirmNewPassword == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "userId, oldPassword, newPassword, and confirmNewPassword are required"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	if newPassword != confirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "New password and confirm new password do not match"})
	}

	// Fetch the user's current password hash from the database
	// res, err := models.GetUser(userId)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid userId"})
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching user"})
	// }

	// user := res.Data.(User)

	// Verify the old password
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Old password is incorrect"})
	// }

	// Update to the new password
	result, err := models.UpdatePassword(userId, newPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateIsAdmin(c echo.Context) error {
	userIdStr := c.FormValue("userId")
	isAdminStr := c.FormValue("isAdmin")

	if userIdStr == "" || isAdminStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "userId and isAdmin are required"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	isAdmin, err := strconv.ParseBool(isAdminStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid isAdmin value"})
	}

	result, err := models.UpdateIsAdmin(userId, isAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateWhatsappNumber(c echo.Context) error {
	userIdStr := c.FormValue("userId")
	whatsappNumber := c.FormValue("whatsappNumber")
	password := c.FormValue("password")

	if userIdStr == "" || whatsappNumber == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "userId and whatsappNumber are required"})
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid userId"})
	}

	// Fetch the user's current password hash from the database
	user, err := models.GetUser(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid userId"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching user"})
	}

	if password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Data.(User).Password), []byte(password))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid password"})
		}
	}

	result, err := models.UpdateWhatsappNumber(userId, whatsappNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
