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

type Store struct {
	StoreId        int
	UserId         int
	Name           string
	PhotoProfile   string
	WhatsappNumber string
	ClosedAt       string
}

func GetStoreIdByUserId(userId int) (Response, error) {
	var res Response
	var storeId int

	con := db.CreateCon()

	sqlStatement := `
		SELECT s.store_id
		FROM "user" u
		JOIN "store" s ON u.user_id = s.user_id
		WHERE u.user_id = $1;
	`
	err := con.QueryRow(sqlStatement, userId).Scan(&storeId)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "No store found for this user"
			return res, nil
		}
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Store ID retrieved successfully"
	res.Data = map[string]int{"store_id": storeId}

	return res, nil
}

func GetStore(storeId int) (Response, error) {
	var store Store
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	checkClosedStatement := `SELECT closed_at FROM store WHERE store_id = $1;`
	var closedAt sql.NullTime
	err := con.QueryRow(checkClosedStatement, storeId).Scan(&closedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "Store not found"
			res.Data = nil
			return res, nil
		}
		return res, err
	}

	if closedAt.Valid {
		res.Success = false
		res.Status = http.StatusNotFound
		res.Message = "Store has been closed"
		res.Data = nil
		return res, nil
	}

	sqlStatement := `
        SELECT s.store_id, s.name AS store_name, s.whatsapp_number, s.photo_profile, 
               p.product_id, p.name AS product_name, p.photo AS product_photo, 
               b.batch_id, b.price
        FROM store s
        LEFT JOIN product p ON s.store_id = p.store_id
        LEFT JOIN batches b ON p.product_id = b.product_id
        WHERE s.store_id = $1 AND p.deleted_at IS NULL;
    `
	rows, err := con.Query(sqlStatement, storeId)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	var productList []map[string]interface{}
	storeData := map[string]interface{}{}

	for rows.Next() {
		var productId *int
		var batchId *int
		var productName *string
		var productPhoto *string
		var productPrice *float64

		err := rows.Scan(&store.StoreId, &store.Name, &store.WhatsappNumber, &store.PhotoProfile,
			&productId, &productName, &productPhoto, &batchId, &productPrice)
		if err != nil {
			return res, err
		}

		if productId != nil && batchId != nil && productName != nil && productPrice != nil {
			productList = append(productList, map[string]interface{}{
				"product_id": *productId,
				"batch_id":   *batchId,
				"name":       *productName,
				"photo":      *productPhoto,
				"price":      *productPrice,
			})
		}
	}

	storeData["store_id"] = store.StoreId
	storeData["store_name"] = store.Name
	storeData["whatsapp_number"] = store.WhatsappNumber
	storeData["profile_photo"] = store.PhotoProfile

	res.Success = true
	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Profile '%s' successfully retrieved", store.Name)
	res.Data = map[string]interface{}{
		"store":    storeData,
		"products": productList,
	}

	return res, nil
}

func GetMyStore(storeId int) (Response, error) {
	var store Store
	var res Response
	var closedAt sql.NullTime

	con := db.CreateCon()
	// defer con.Close()

	sqlStatement := `
        SELECT store_id, photo_profile, name, whatsapp_number, closed_at
        FROM "store" 
        WHERE store_id = $1;
    `
	row := con.QueryRow(sqlStatement, storeId)

	err := row.Scan(&store.StoreId, &store.PhotoProfile, &store.Name, &store.WhatsappNumber, &closedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "Store not found"
			res.Data = nil
			return res, nil
		}
		return res, err
	}

	storeStatus := "open"
	if closedAt.Valid {
		storeStatus = "closed"
	}

	sqlProductCounts := `
		SELECT
			COUNT(*) AS jumlah_produk
		FROM "product"
		WHERE store_id = $1;`

	var jumlahProduk int

	err = con.QueryRow(sqlProductCounts, storeId).Scan(
		&jumlahProduk,
	)
	if err != nil {
		return res, err
	}

	sqlOrderCounts := `
		SELECT
			COUNT(*) AS total_order
		FROM "orders" o
		INNER JOIN "batches" b ON o.batch_id = b.batch_id
		INNER JOIN "product" p ON b.product_id = p.product_id
		WHERE p.store_id = $1;
	`

	var totalOrder int

	err = con.QueryRow(sqlOrderCounts, storeId).Scan(&totalOrder)
	if err != nil {
		return res, err
	}

	sqlPesananHabis := `
		SELECT
		COUNT(*) AS total_pesanan_habis
		FROM "batches" b
		INNER JOIN "product" p ON b.product_id = p.product_id
		WHERE p.store_id = $1 AND b.stock = 0;
	`

	var totalPesananHabis int

	err = con.QueryRow(sqlPesananHabis, storeId).Scan(&totalPesananHabis)
	if err != nil {
		return res, err
	}

	currentTime := time.Now()

	sqlPesananDitutup := `
		SELECT
		COUNT(*) AS total_pesanan_ditutup
		FROM "batches" b
		INNER JOIN "product" p ON b.product_id = p.product_id
		WHERE p.store_id = $1 AND b.close_order_time < $2;
	`

	var totalPesananDitutup int

	err = con.QueryRow(sqlPesananDitutup, storeId, currentTime).Scan(&totalPesananDitutup)
	if err != nil {
		return res, err
	}

	sqlBatchTersedia := `
		SELECT
			COUNT(*) AS total_batch_tersedia
		FROM "batches"
		WHERE product_id IN (SELECT product_id FROM "product" WHERE store_id = $1) AND stock != 0 AND close_order_time > $2;
	`

	var totalBatchTersedia int

	err = con.QueryRow(sqlBatchTersedia, storeId, currentTime).Scan(&totalBatchTersedia)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Profile '%s' successfully retrieved", store.Name)
	res.Data = map[string]interface{}{
		"store": map[string]interface{}{
			"storeId":        store.StoreId,
			"name":           store.Name,
			"photoProfile":   store.PhotoProfile,
			"whatsappNumber": store.WhatsappNumber,
			"status":         storeStatus,
		},
		"products": map[string]int{
			"jumlahProduk": jumlahProduk},
		"orders": map[string]int{
			"totalBatchTersedia":  totalBatchTersedia,
			"totalOrder":          totalOrder,
			"totalPesananHabis":   totalPesananHabis,
			"totalPesananDitutup": totalPesananDitutup,
		},
	}

	return res, nil
}

func CreateStore(userId int, name string, photoProfile string, whatsappNumber string) (Response, error) {
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	var existingStoreId int
	var closedAt sql.NullTime
	checkStoreQuery := `
        SELECT store_id, closed_at
        FROM "store"
        WHERE user_id = $1; 
    `
	err := con.QueryRow(checkStoreQuery, userId).Scan(&existingStoreId, &closedAt)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	if existingStoreId != 0 && !closedAt.Valid {
		res.Success = false
		res.Status = http.StatusConflict
		res.Message = "User already has an active store"
		return res, nil
	}

	if existingStoreId != 0 && closedAt.Valid {
		updateQuery := `
            UPDATE "store"
            SET closed_at = NULL
            WHERE store_id = $1;
        `
		stmt, err := con.Prepare(updateQuery)
		if err != nil {
			return res, err
		}
		defer stmt.Close()

		_, err = stmt.Exec(existingStoreId)
		if err != nil {
			return res, err
		}

		res.Success = true
		res.Status = http.StatusOK
		res.Message = "Success to restore store"
		res.Data = map[string]int{"RestoreStoreId": existingStoreId}
		return res, nil
	}

	var existingStoreNameId int
	var existingStoreClosedAt sql.NullTime
	checkStoreNameQuery := `
        SELECT store_id, closed_at
        FROM "store"
        WHERE name = $1; 
    `
	err = con.QueryRow(checkStoreNameQuery, name).Scan(&existingStoreNameId, &existingStoreClosedAt)
	if err != nil && err != sql.ErrNoRows {
		return res, err
	}

	if existingStoreNameId != 0 && !existingStoreClosedAt.Valid {
		res.Success = false
		res.Status = http.StatusConflict
		res.Message = "A store with the same name already exists"
		return res, nil
	}

	sqlStatement := `
        INSERT INTO "store" (user_id, name, photo_profile, whatsapp_number) 
        VALUES($1, $2, $3, $4) 
        RETURNING store_id;
    `

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(userId, name, photoProfile, whatsappNumber).Scan(&id)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Success to create new store: %s", name)
	res.Data = map[string]any{"name": name, "photoProfile": photoProfile, "whatsappNumber": whatsappNumber, "LastStoreId": id}

	return res, nil
}

func UpdateStore(storeId int, storeName string, photoProfile string, whatsappNumber string) (Response, error) {
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	var updateValues []interface{}
	var sqlValues []string

	columns := []struct {
		name  string
		value string
	}{
		{"name", storeName},
		{"photo_profile", photoProfile},
		{"whatsapp_number", whatsappNumber},
	}

	// Prepare SQL set clauses and parameter values
	for _, col := range columns {
		if col.value != "" {
			sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
			updateValues = append(updateValues, col.value)
		}
	}

	// Check if there are values to update
	if len(sqlValues) == 0 {
		res.Success = false
		res.Status = http.StatusBadRequest
		res.Message = "No data to update"
		return res, fmt.Errorf("no data to update")
	}

	// Add storeId to parameter values
	updateValues = append(updateValues, storeId)

	sqlStatement := "UPDATE \"store\" SET " + strings.Join(sqlValues, ", ") + " WHERE store_id = $" + strconv.Itoa(len(updateValues)) + ";"

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
	res.Message = "Success update store!"
	res.Data = map[string]int64{"rowsAffected": rowsAffected}

	return res, nil
}

func OpenStore(storeId int) (Response, error) {
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	var closedAt sql.NullTime
	checkSql := `
        SELECT closed_at FROM "store" WHERE store_id = $1;
    `

	err := con.QueryRow(checkSql, storeId).Scan(&closedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "Store not found!"
			return res, nil
		}
		return res, err
	}

	if !closedAt.Valid {
		res.Success = false
		res.Status = http.StatusConflict
		res.Message = "Store is already open!"
		return res, nil
	}

	sqlStatement := `
        UPDATE "store" SET closed_at = NULL 
        WHERE store_id = $1;
    `

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(storeId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to open the store!"
	res.Data = map[string]int64{"rows": rowsAffected}

	return res, nil
}

func CloseStore(storeId int) (Response, error) {
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	var closedAt sql.NullTime
	checkSql := `
        SELECT closed_at FROM "store" WHERE store_id = $1;
    `

	err := con.QueryRow(checkSql, storeId).Scan(&closedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			res.Status = http.StatusNotFound
			res.Message = "Store not found!"
			return res, nil
		}
		return res, err
	}

	if closedAt.Valid {
		res.Success = false
		res.Status = http.StatusConflict
		res.Message = "Store is already closed!"
		return res, nil
	}

	closedAt = sql.NullTime{Time: time.Now(), Valid: true}
	sqlStatement := `
        UPDATE "store" SET closed_at = $1 
        WHERE store_id = $2;
    `

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(closedAt, storeId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to close the store!"
	res.Data = map[string]int64{"rows": rowsAffected}

	return res, nil
}

func CheckUserHasStore(userId int) (map[string]interface{}, error) {
	con := db.CreateCon()
	// defer con.Close()

	var storeId int
	sqlStatement := `SELECT store_id FROM "store" WHERE user_id = $1 LIMIT 1`
	err := con.QueryRow(sqlStatement, userId).Scan(&storeId)

	if err != nil {
		if err == sql.ErrNoRows {
			return map[string]interface{}{"hasStore": false, "storeId": int64(0)}, nil // User has no store
		}
		return nil, err
	}

	return map[string]interface{}{"hasStore": true, "storeId": storeId}, nil // User has a store
}
