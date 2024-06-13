// file: auth_service/controller/auth.go
package auth_controller

import (
	"backend/models"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Username             string `json:"username"`
	WhatsappNumber       string `json:"whatsapp_number"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func Register(c echo.Context) error {
	var creds Credentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	fmt.Println("Email:", creds.Password)
	fmt.Println("Username:", creds.Username)
	fmt.Println("WhatsappNumber:", creds.WhatsappNumber)

	// Check if password and password confirmation match
	if creds.Password != creds.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Passwords do not match"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
	}

	res, err := models.CreateUser(creds.Email, creds.Username, creds.WhatsappNumber, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	var creds Credentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	fmt.Println(creds.Email)
	res, err := models.GetUserByEmail(creds.Email)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching user"})
	}

	user := res.Data.(models.User)
	fmt.Println(user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": "true",
		"status":  http.StatusOK,
		"message": "Success to login!",
		"token":   tokenString,
	})
}

func ValidateToken(c echo.Context) error {
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

	res, err := models.GetUser(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching user"})
	}

	return c.JSON(http.StatusOK, res)
}

func RefreshToken(c echo.Context) error {
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

	// Check if token is about to expire
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Minute {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Token is not close to expiration"})
	}

	// Create new token with extended expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		c.Set("user", token)
		return next(c)
	}
}
