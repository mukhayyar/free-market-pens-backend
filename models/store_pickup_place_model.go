package models

import (
	"backend/db"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

type StorePickupPlace struct {
	StorePickupPlaceId        int
	StoreId          		  int
	Name             		  string
}

func GetAllStorePickupPlace(storeId int) (Response, error) {
    var pickupPlace StorePickupPlace
    var pickupPlaceList []StorePickupPlace
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT store_pickup_place_id, store_id, name
        FROM store_pickup_place
        WHERE store_id = $1;
    `
    rows, err := con.Query(sqlStatement, storeId)
    if err != nil {
        return res, err
    }
    defer rows.Close()

    for rows.Next(){
		err = rows.Scan(&pickupPlace.StorePickupPlaceId, &pickupPlace.StoreId, &pickupPlace.Name)
		if err != nil{
			return res, err
		}
		pickupPlaceList = append(pickupPlaceList, pickupPlace)
	}

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to list the place of pickup store"
    res.Data = pickupPlaceList

    return res, nil
}

func CreateStorePickupPlace(storeId int, name string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := `
		INSERT INTO "store_pickup_place" (store_id, name) 
		VALUES($1, $2) 
		RETURNING store_pickup_place_id;
	`
	
    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(storeId, name).Scan(&id)
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok {
            if pqErr.Code.Name() == "unique_store_pickup_place" {
                return res, fmt.Errorf("a place with the same name already exists")
            }
        }
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add store pickup place!"
    res.Data = map[string]int64{"LastStorePickupPlaceId": id}

    return res, nil
}