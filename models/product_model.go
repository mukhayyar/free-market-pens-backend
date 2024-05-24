package models

import (
	"backend/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ProductId        int
	StoreId          int
	Name             string
	Photo            string
    CategoryId       int
    Description      string
    DeletedAt        string
}

func GetMyProductDetail(productId int) (Response, error) {
    var product Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT product_id, photo, name, description
        FROM "product"
        WHERE product_id = $1 AND deleted_at IS NULL;
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
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	var storeName string
	storeNameQuery := `SELECT name FROM "store" WHERE store_id = $1`
	err := con.QueryRow(storeNameQuery, storeId).Scan(&storeName)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "Store not found"
			return res, nil
		}
		return res, err
	}

	sqlStatement := `
        SELECT p.product_id, p.store_id, p.name, p.photo, b.price
        FROM product p
        LEFT JOIN batches b ON p.product_id = b.product_id
        WHERE p.store_id = $1 AND p.deleted_at IS NULL;
    `
	rows, err := con.Query(sqlStatement, storeId)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	var productList []map[string]interface{}
	for rows.Next() {
		var product Product
		var price sql.NullFloat64

		err = rows.Scan(&product.ProductId, &product.StoreId, &product.Name, &product.Photo, &price)
		if err != nil {
			return res, err
		}

		var priceValue interface{}
		if price.Valid {
			priceValue = price.Float64
		} else {
			priceValue = ""
		}

		productData := map[string]interface{}{
			"id":    product.ProductId,
			"name":  product.Name,
			"photo": product.Photo,
			"price": priceValue,
		}
		productList = append(productList, productData)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Success to show list product of store: %s", storeName)
	res.Data = map[string]interface{}{
		"storeId":  storeId,
		"products": productList,
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

    // Kalo produk di toko yang sama ada
    var existingProductId int64
    var deletedAt sql.NullTime
    checkProductQuery := `
        SELECT product_id, deleted_at
        FROM "product"
        WHERE store_id = $1 AND name = $2;
    `
    err := con.QueryRow(checkProductQuery, storeId, name).Scan(&existingProductId, &deletedAt)
    if err != nil && err != sql.ErrNoRows {
        return res, err
    }

    // Nama produknya sama dan tidak dihapus
    if existingProductId != 0 && !deletedAt.Valid {
        res.Success = false
        res.Status = http.StatusConflict
        res.Message = "A product with the same name already exists"
        return res, nil
    }

    // Nama produknya sama, tapi udah dihapus
    if existingProductId != 0 && deletedAt.Valid {
        updateQuery := `
            UPDATE "product"
            SET deleted_at = NULL
            WHERE product_id = $1 AND store_id = $2;
        `
        stmt, err := con.Prepare(updateQuery)
        if err != nil {
            return res, err
        }
        defer stmt.Close()

        _, err = stmt.Exec(existingProductId, storeId)
        if err != nil {
            return res, err
        }

        res.Success = true
        res.Status = http.StatusOK
        res.Message = "Success to update and restore product!"
        res.Data = map[string]int64{"RestoredProductId": existingProductId}
        return res, nil
    }

    // Produk tidak ada sebelumnya
    sqlStatement := `
        INSERT INTO "product" (store_id, photo, name, description)
        VALUES ($1, $2, $3, $4)
        RETURNING product_id;
    `
    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var newProductId int64
    err = stmt.QueryRow(storeId, photo, name, description).Scan(&newProductId)
    if err != nil {
        return res, err
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add product!"
    res.Data = map[string]int64{"LastProductId": newProductId}

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

    if len(sqlValues) == 0 {
        res.Success = false
        res.Status = http.StatusBadRequest
        res.Message = "No data to update"
        return res, fmt.Errorf("no data to update")
    }

    updateValues = append(updateValues, productId)

    sqlStatement := "UPDATE \"product\" SET " + strings.Join(sqlValues, ", ") + " WHERE product_id = $" + strconv.Itoa(len(updateValues)) + ";"

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

    if rowsAffected == 0 {
		res.Status = http.StatusNotFound
		res.Message = "No product found with the given id"
		res.Success = false
		return res, fmt.Errorf("no product found with the given id")
	}

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success update product!"
    res.Data = map[string]int64{"rowsAffected   ": rowsAffected}

    return res, nil
}

func DeleteProduct(productId int) (Response, error) {
	var res Response

	con := db.CreateCon()

    deletedAt := time.Now()
	sqlStatement := `
        UPDATE "product" SET deleted_at = $1 
        WHERE product_id = $2;
    `

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	} 
	
	result, err := stmt.Exec(deletedAt, productId)
	if err != nil {
		return res, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

    res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success delete product!"
	res.Data = map[string]int64{"rows": rowsAffected}

	return res, nil
}