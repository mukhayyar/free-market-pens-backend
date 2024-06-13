package models

import (
	"backend/db"
	"database/sql"
	"net/http"
)

type Transaction struct {
	TransactionID              int     `json:"transaction_id"`
	UserID                     int     `json:"user_id"`
	TransactionDate            string  `json:"transaction_date"`
	TotalPayment               float64 `json:"total_payment"`
	Quantity                   int     `json:"quantity"`
	TransactionStatus          string  `json:"transaction_status"`
	CancelledTransactionDate   string  `json:"cancelled_transaction_date,omitempty"`
	CancelledTransactionReason string  `json:"cancelled_transaction_reason,omitempty"`
}

func GetTransaction(transactionID int) (Response, error) {
	var transaction Transaction
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM transactions WHERE transaction_id = $1`
	row := con.QueryRow(sqlStatement, transactionID)
	err := row.Scan(&transaction.TransactionID, &transaction.UserID, &transaction.TransactionDate, &transaction.TotalPayment, &transaction.Quantity, &transaction.TransactionStatus, &transaction.CancelledTransactionDate, &transaction.CancelledTransactionReason)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		}
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the transaction data"
	res.Data = transaction

	return res, nil
}

func GetTransactionsByUserID(userID int) (Response, error) {
	var transactions []Transaction
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM transactions WHERE user_id = $1`
	rows, err := con.Query(sqlStatement, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.TransactionID, &transaction.UserID, &transaction.TransactionDate, &transaction.TotalPayment, &transaction.Quantity, &transaction.TransactionStatus, &transaction.CancelledTransactionDate, &transaction.CancelledTransactionReason)
		if err != nil {
			return res, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the transactions data"
	res.Data = transactions

	return res, nil
}

func GetTransactionsByProductID(productID int) (Response, error) {
	var transactions []Transaction
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM transactions WHERE product_id = $1`
	rows, err := con.Query(sqlStatement, productID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.TransactionID, &transaction.UserID, &transaction.TransactionDate, &transaction.TotalPayment, &transaction.Quantity, &transaction.TransactionStatus, &transaction.CancelledTransactionDate, &transaction.CancelledTransactionReason)
		if err != nil {
			return res, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the transactions data"
	res.Data = transactions

	return res, nil
}

func GetTransactionsByBatchID(batchID int) (Response, error) {
	var transactions []Transaction
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM transactions WHERE product_id = $1`
	rows, err := con.Query(sqlStatement, batchID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.TransactionID, &transaction.UserID, &transaction.TransactionDate, &transaction.TotalPayment, &transaction.Quantity, &transaction.TransactionStatus, &transaction.CancelledTransactionDate, &transaction.CancelledTransactionReason)
		if err != nil {
			return res, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the transactions data"
	res.Data = transactions

	return res, nil
}

func CreateTransaction(userID int, transactionDate string, totalPayment float64, quantity int, transactionStatus, cancelledTransactionDate, cancelledTransactionReason string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `
		INSERT INTO transactions (user_id,  transaction_date, total_payment, quantity, transaction_status, cancelled_transaction_date, cancelled_transaction_reason) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING transaction_id;
	`
	var transactionID int
	err := con.QueryRow(sqlStatement, userID, transactionDate, totalPayment, quantity, transactionStatus, cancelledTransactionDate, cancelledTransactionReason).Scan(&transactionID)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to create transaction"
	res.Data = map[string]int{"transaction_id": transactionID}

	return res, nil
}

func UpdateTransaction(transactionID int, userID int, transactionDate string, totalPayment float64, quantity int, transactionStatus, cancelledTransactionDate, cancelledTransactionReason string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `
		UPDATE transactions
		SET user_id = $1, product_id = $2, batch_id = $3, transaction_date = $4, total_payment = $5, quantity = $6, transaction_status = $7, cancelled_transaction_date = $8, cancelled_transaction_reason = $9
		WHERE transaction_id = $10;
	`
	result, err := con.Exec(sqlStatement, userID, transactionDate, totalPayment, quantity, transactionStatus, cancelledTransactionDate, cancelledTransactionReason, transactionID)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update transaction"
	res.Data = map[string]int64{"rows_affected": rowsAffected}

	return res, nil
}

func DeleteTransaction(transactionID int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `DELETE FROM transactions WHERE transaction_id = $1`
	result, err := con.Exec(sqlStatement, transactionID)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to delete transaction"
	res.Data = map[string]int64{"rows_affected": rowsAffected}

	return res, nil
}

func UpdateTransactionStatus(transactionID int, status string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `
		UPDATE transactions
		SET transaction_status = $1
		WHERE transaction_id = $2;
	`
	result, err := con.Exec(sqlStatement, status, transactionID)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update transaction status"
	res.Data = map[string]int64{"rows_affected": rowsAffected}

	return res, nil
}
