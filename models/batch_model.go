package models

import (
	"backend/db"
	"net/http"
)

type Batch struct {
	BatchId        int
	ProductId      int
	PickupPlaceId  int
	Stock          int
	CloseOrderDate string
	CloseOrderTime string
	PickupDate     string
	PickupTime     string
}

func GetBatchById()  {
    
}

func CreateBatch(productId int, pickupPlaceId int, stock int, pickupDate string, pickupTime string, closeOrderDate string, closeOrderTime string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := `
		INSERT INTO "batches" (product_id, store_pickup_place_id, stock, close_order_date, close_order_time, pickup_date, pickup_time) 
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING batch_id;
	`

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(productId, pickupPlaceId, stock, pickupDate, pickupTime, closeOrderDate, closeOrderTime).Scan(&id)
    if err != nil {
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add batch!"
    res.Data = map[string]int64{"LastBatchId": id}

    return res, nil
}