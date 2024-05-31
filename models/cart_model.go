package models

import (
	"backend/db"
	"database/sql"
	"net/http"
)

// Cart represents a cart item
type Cart struct {
	CartID    int `json:"cart_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// GetCart retrieves a cart item by ID
func GetCart(cartID int) (Response, error) {
	var cart Cart
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM cart WHERE cart_id = $1`
	row := con.QueryRow(sqlStatement, cartID)
	err := row.Scan(&cart.CartID, &cart.UserID, &cart.ProductID, &cart.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		}
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the cart item"
	res.Data = cart

	return res, nil
}

// GetCartsByUserID retrieves all cart items for a user
func GetCartsByUserID(userID int) (Response, error) {
	var carts []Cart
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM cart WHERE user_id = $1`
	rows, err := con.Query(sqlStatement, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var cart Cart
		err := rows.Scan(&cart.CartID, &cart.UserID, &cart.ProductID, &cart.Quantity)
		if err != nil {
			return res, err
		}
		carts = append(carts, cart)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the cart items"
	res.Data = carts

	return res, nil
}

// CreateCart adds a new item to the cart
func CreateCart(userID, productID, quantity int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `
		INSERT INTO cart (user_id, product_id, quantity) 
		VALUES ($1, $2, $3)
		RETURNING cart_id;
	`
	var cartID int
	err := con.QueryRow(sqlStatement, userID, productID, quantity).Scan(&cartID)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to add item to cart"
	res.Data = map[string]int{"cart_id": cartID}

	return res, nil
}

// UpdateCart updates the quantity of a cart item
func UpdateCart(cartID, quantity int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `
		UPDATE cart
		SET quantity = $1
		WHERE cart_id = $2;
	`
	result, err := con.Exec(sqlStatement, quantity, cartID)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update cart item"
	res.Data = map[string]int64{"rows_affected": rowsAffected}

	return res, nil
}

// DeleteCart removes an item from the cart
func DeleteCart(cartID int) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `DELETE FROM cart WHERE cart_id = $1`
	result, err := con.Exec(sqlStatement, cartID)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to delete cart item"
	res.Data = map[string]int64{"rows_affected": rowsAffected}

	return res, nil
}
