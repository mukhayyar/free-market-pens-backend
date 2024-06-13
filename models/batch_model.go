package models

import (
	"backend/db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Batch struct {
	BatchId        int
	ProductId      int
	PickupPlaceId  int
	Stock          int
	Price          float64
	CloseOrderTime string
	PickupTime     string
}

func GetLastBatch(productId int) (Response, error) {
	var batch Batch
	var storePickupPlace StorePickupPlace
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	// Count total batches for the product to determine the batch number
	countSqlStatement := `
        SELECT COUNT(*)
        FROM batches
        WHERE product_id = $1;
    `
	var totalBatches int
	err := con.QueryRow(countSqlStatement, productId).Scan(&totalBatches)
	if err != nil {
		return res, err
	}

	if totalBatches == 0 {
		// No batches found for this product
		res.Success = true
		res.Status = http.StatusOK
		res.Message = "No batches found for this product"
		res.Data = nil
		return res, nil
	}

	// Fetch the latest batch
	sqlStatement := `
        SELECT b.batch_id, b.stock, b.price, b.close_order_time, b.pickup_time, spp.name 
        FROM batches b
        LEFT JOIN store_pickup_place spp ON b.store_pickup_place_id = spp.store_pickup_place_id
        WHERE b.product_id = $1
        ORDER BY b.batch_id DESC
        LIMIT 1;
    `
	err = con.QueryRow(sqlStatement, productId).Scan(&batch.BatchId, &batch.Stock, &batch.Price, &batch.CloseOrderTime, &batch.PickupTime, &storePickupPlace.Name)
	if err != nil {
		return res, err
	}

	batchData := map[string]interface{}{
		"id":           batch.BatchId,
		"batch_number": totalBatches,
		"stock":        batch.Stock,
		"price":        batch.Price,
		"close_order": map[string]interface{}{
			"date": strings.Split(batch.CloseOrderTime, "T")[0],
			"time": strings.Split(batch.CloseOrderTime, "T")[1],
		},
		"pickup": map[string]interface{}{
			"date":  strings.Split(batch.PickupTime, "T")[0],
			"time":  strings.Split(batch.PickupTime, "T")[1],
			"place": storePickupPlace.Name,
		},
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to retrieve the last batch data"
	res.Data = batchData

	return res, nil
}

func GetAllBatch(productId int) (Response, error) {
	var batch Batch
	var storePickupPlace StorePickupPlace
	var batchList []Batch
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	sqlStatement := `
        SELECT b.batch_id, b.stock, b.price, b.close_order_time, b.pickup_time, spp.name 
        FROM batches b
        LEFT JOIN store_pickup_place spp ON b.store_pickup_place_id = spp.store_pickup_place_id
        WHERE b.product_id = $1;
    `

	rows, err := con.Query(sqlStatement, productId)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&batch.BatchId, &batch.Stock, &batch.Price, &batch.CloseOrderTime, &batch.PickupTime, &storePickupPlace.Name)
		if err != nil {
			return res, err
		}
		batchList = append(batchList, batch)
	}

	var batches []map[string]interface{}
	for _, batch := range batchList {
		batchData := map[string]interface{}{
			"id":    batch.BatchId,
			"stock": batch.Stock,
			"price": batch.Price,
			"close_order": map[string]interface{}{
				"date": strings.Split(batch.CloseOrderTime, "T")[0],
				"time": strings.Split(batch.CloseOrderTime, "T")[1],
			},
			"pickup": map[string]interface{}{
				"date":  strings.Split(batch.PickupTime, "T")[0],
				"time":  strings.Split(batch.PickupTime, "T")[1],
				"place": storePickupPlace.Name,
			},
		}
		batches = append(batches, batchData)
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to list batch data"
	res.Data = map[string]interface{}{
		"batchList": batches,
	}

	return res, nil

}

func CreateBatch(productId int, pickupPlaceId int, stock int, price float64, pickupTime string, closeOrderTime string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
		INSERT INTO "batches" (product_id, store_pickup_place_id, stock, price, pickup_time, close_order_time) 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING batch_id;
	`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(productId, pickupPlaceId, stock, price, pickupTime, closeOrderTime).Scan(&id)
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to add batch!"
	res.Data = map[string]int64{"LastBatchId": id}

	return res, nil
}

func UpdateBatch(batchId int, pickupPlaceId int, stock int, price float64, pickupTime string, closeOrderTime string) (Response, error) {
	var res Response

	con := db.CreateCon()

	var updateValues []interface{}
	var sqlValues []string

	columns := []struct {
		name  string
		value any
	}{
		{"store_pickup_place_id", pickupPlaceId},
		{"stock", stock},
		{"price", price},
		{"pickup_time", pickupTime},
		{"close_order_time", closeOrderTime},
	}

	for _, col := range columns {
		switch v := col.value.(type) {
		case int:
			if v != 0 {
				sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
				updateValues = append(updateValues, col.value)
			}
		case float64:
			if v != 0.0 {
				sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
				updateValues = append(updateValues, col.value)
			}
		case string:
			if v != "" {
				sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
				updateValues = append(updateValues, col.value)
			}
		}
	}

	if len(sqlValues) == 0 {
		res.Success = false
		res.Status = http.StatusBadRequest
		res.Message = "No data to update"
		return res, fmt.Errorf("no data to update")
	}

	updateValues = append(updateValues, batchId)

	sqlStatement := "UPDATE \"batches\" SET " + strings.Join(sqlValues, ", ") + " WHERE batch_id = $" + strconv.Itoa(len(updateValues)) + ";"

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
		res.Message = "No batch found with the given id"
		res.Success = false
		return res, fmt.Errorf("no batch found with the given id")
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success update batch!"
	res.Data = map[string]int64{"rowsAffected   ": rowsAffected}

	return res, nil
}
