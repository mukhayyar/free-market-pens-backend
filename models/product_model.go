package models

import (
	"backend/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type Product struct {
	ProductId        int
	StoreId          int
	Name             string
	Photo            string
    CategoryId       int
    Description      string
}

func GetMyProductDetail(productId int) (Response, error) {
    var product Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT product_id, photo, name, description
        FROM "product"
        WHERE product_id = $1;
    `
    row := con.QueryRow(sqlStatement, productId)

    err := row.Scan(&product.ProductId, &product.Photo, &product.Name, &product.Description)
    if err != nil {
        if err == sql.ErrNoRows {
            return res, err
        }
        return res, err
    }    

    res.Success = true
    res.Status = http.StatusOK
    res.Message = fmt.Sprintf("Product detail of '%s' successfully retrieved", product.Name)
    res.Data = map[string]interface{}{
        "product": map[string]interface{}{
            "productId":      product.ProductId,
            "name":           product.Name,
            "productPhoto":   product.Photo,
            "description": product.Description,
        },
    }

    return res, nil
}

func GetAllMyProduct(storeId int) (Response, error) {
    var product Product
    var productList []Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT product_id, store_id, name
        FROM product
        WHERE store_id = $1;
    `
    rows, err := con.Query(sqlStatement, storeId)
    if err != nil {
        return res, err
    }
    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&product.ProductId, &product.StoreId, &product.Name)
        if err != nil {
            return res, err
        }
        productList = append(productList, product)
    }

    var products []map[string]interface{}
    for _, prod := range productList {
        productData := map[string]interface{}{
            "id":    prod.ProductId,
            "name":  prod.Name,
        }
        products = append(products, productData)
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Products Home successfully retrieved"
    res.Data = map[string]interface{}{
        "storeId":   storeId,
        "products": products,
    }

    return res, nil
}

func GetAllProduct(closeOrderDate string, pickupDate string) (Response, error) {
    var product Product
    var batch Batch
    var productList []Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT p.product_id, p.photo, p.name, b.price
        FROM batches b
        LEFT JOIN product p ON b.product_id = p.product_id
    `
    
    var conditions []string
    var params []interface{}
    if closeOrderDate != "" {
        conditions = append(conditions, "DATE(b.close_order_time) = $1"+strconv.Itoa(len(params)+1))
        params = append(params, closeOrderDate)
    }
    if pickupDate != "" {
        conditions = append(conditions, "DATE(b.pickup_time) = $2"+strconv.Itoa(len(params)+1))
        params = append(params, pickupDate)
    }

    if len(conditions) > 0 {
        sqlStatement += " WHERE " + strings.Join(conditions, " AND ") + ";"
    }

    rows, err := con.Query(sqlStatement, params...)
    if err != nil {
        return res, err
    }
    defer rows.Close()

    for rows.Next() {
        err = rows.Scan(&product.ProductId, &product.Photo, &product.Name, &batch.Price)
        if err != nil {
            return res, err
        }
        productList = append(productList, product)
    }

    var products []map[string]interface{}
    for _, prod := range productList {
        productData := map[string]interface{}{
            "id":    prod.ProductId,
            "photo": prod.Photo,
            "name":  prod.Name,
            "price": batch.Price,
        }
        products = append(products, productData)
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Products Home successfully retrieved"
    res.Data = map[string]interface{}{
        "filter": map[string]interface{}{
            "close_order": closeOrderDate,
            "pickup_date": pickupDate,
        },
        "products": products,
    }

    return res, nil
}

func CreateProduct(storeId int, photo string, name string, description string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := `
        INSERT INTO "product" (store_id, photo, name, description) 
        VALUES($1, $2, $3, $4) 
        RETURNING product_id;
    `

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(storeId, photo, name, description).Scan(&id)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok {
            if pqErr.Code.Name() == "unique_product_name" {
                return res, fmt.Errorf("a product with the same name already exists")
            }
        }
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add product!"
    res.Data = map[string]int64{"LastProductId": id}

    return res, nil
}

func UpdateProduct(productId int, photo string, name string, description string) (Response, error) {
    var res Response

    con := db.CreateCon()

    var updateValues []interface{}
    var sqlValues []string

    columns := []struct {
        name  string
        value string
    }{
        {"photo", photo},
        {"name", name},
        {"description", description},
    }

    for _, col := range columns {
        if col.value != "" {
            sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
            updateValues = append(updateValues, col.value)
        }
    }

    sqlStatement := "UPDATE \"product\" SET " + strings.Join(sqlValues, ", ") + " WHERE product_id = $1" + strconv.Itoa(len(updateValues)+1) + ";"
    updateValues = append(updateValues, productId)

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(updateValues...)
    if err != nil {
        return res, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return res, err
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success update product!"
    res.Data = map[string]int64{"rowsAffected   ": rowsAffected}

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