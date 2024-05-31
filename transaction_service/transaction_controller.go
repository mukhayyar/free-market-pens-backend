package transaction_controller

import (
	"backend/models"
	"database/sql"
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

func GetTransactionByID(c echo.Context) error {
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	res, err := models.GetTransaction(transactionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Transaction not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func GetTransactionsByUserID(c echo.Context) error {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	res, err := models.GetTransactionsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func GetTransactionsByProductID(c echo.Context) error {
	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid product ID"})
	}

	res, err := models.GetTransactionsByProductID(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func GetTransactionsByBatchID(c echo.Context) error {
	batchIDStr := c.Param("batchID")
	batchID, err := strconv.Atoi(batchIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid batch ID"})
	}

	res, err := models.GetTransactionsByBatchID(batchID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func CreateTransaction(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	res, err := models.CreateTransaction(transaction.UserID, transaction.ProductID, transaction.BatchID, transaction.TransactionDate, transaction.TotalPayment, transaction.Quantity, transaction.TransactionStatus, transaction.CancelledTransactionDate, transaction.CancelledTransactionReason)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateTransaction(c echo.Context) error {
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	res, err := models.UpdateTransaction(transactionID, transaction.UserID, transaction.ProductID, transaction.BatchID, transaction.TransactionDate, transaction.TotalPayment, transaction.Quantity, transaction.TransactionStatus, transaction.CancelledTransactionDate, transaction.CancelledTransactionReason)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteTransaction(c echo.Context) error {
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	res, err := models.DeleteTransaction(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateTransactionStatusPaid(c echo.Context) error {
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	res, err := models.UpdateTransactionStatus(transactionID, "PAID")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateTransactionStatusUnpaid(c echo.Context) error {
	transactionIDStr := c.Param("transactionID")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	res, err := models.UpdateTransactionStatus(transactionID, "UNPAID")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
