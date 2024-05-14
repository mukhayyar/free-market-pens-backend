package models

import (
	"backend/db"
	"net/http"
)

type Product struct {
	ProductId        int
	StoreId          int
	Name             string
	Photo            string
	Price            float64
    CategoryId       int
    Description      string
}

func GetAllMyProduct(storeId int) (Response, error) {
    var product Product
    var productList []Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT product_id, store_id, name, price, photo
        FROM product
        WHERE store_id = $1;
    `
    rows, err := con.Query(sqlStatement, storeId)
    if err != nil {
        return res, err
    }
    defer rows.Close()

    for rows.Next(){
		err = rows.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Price, &product.Photo)
		if err != nil{
			return res, err
		}
		productList = append(productList, product)
	}

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to list the product of store"
    res.Data = productList

    return res, nil
}

func CreateProduct(storeId string, photo string, name string, description string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := "INSERT INTO \"product\" (store_id, photo, name, description, price) VALUES($1, $2, $3, $4, 0) RETURNING product_id;"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(storeId, photo, name, description).Scan(&id)
    if err != nil {
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add product!"
    res.Data = map[string]int64{"LastProductId": id}

    return res, nil
}

func AddPrice(productId int, price float64) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := `
        UPDATE "product"
        SET price = $2
        WHERE product_id = $1;
    `
    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(stmt)
    if err != nil {
        return res, err
    }

    rowAffected, err := result.LastInsertId()
    if err != nil {
        return res, err
    }

    res.Status = http.StatusOK
    res.Message = "Success add price!"
    res.Data = map[string]int64{"row_affected": rowAffected}

    return res, nil
}

// func DeleteUser(userId int) (Response, error) {
// 	var res Response

// 	con := db.CreateCon()

// 	sqlStatement := "DELETE FROM \"user\" WHERE user_id = $1;"

// 	stmt, err := con.Prepare(sqlStatement)
// 	if err != nil {
// 		return res, err
// 	} 
	
// 	result, err := stmt.Exec(userId)
// 	if err != nil {
// 		return res, err
// 	}
	
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = http.StatusOK
// 	res.Message = "Success Delete User!"
// 	res.Data = map[string]int64{"rows": rowsAffected}

// 	return res, nil
// }