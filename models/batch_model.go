package models

import (
	"backend/db"
	"net/http"
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
    if err != nil{
        return res, err
    }
    defer rows.Close()

    for rows.Next(){
        err = rows.Scan(&batch.BatchId, &batch.Stock, &batch.Price, &batch.CloseOrderTime, &batch.PickupTime, &storePickupPlace.Name)
        if err != nil{
            return res, err
        }
        batchList = append(batchList, batch)
    }

    var batches []map[string]interface{}
    for _, batch := range batchList {
        batchData := map[string]interface{}{
            "id": batch.BatchId,
            "stock": batch.Stock,
            "price": batch.Price,
            "close_order": map[string]interface{}{
                "date": strings.Split(batch.CloseOrderTime, "T")[0],
                "time": strings.Split(batch.CloseOrderTime, "T")[1],
            },
            "pickup": map[string]interface{}{
                "date": strings.Split(batch.PickupTime, "T")[0],
                "time": strings.Split(batch.PickupTime, "T")[1],
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
		INSERT INTO "batches" (product_id, store_pickup_place_id, stock, price, close_order_time, pickup_time) 
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