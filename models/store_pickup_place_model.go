package models

import (
	"backend/db"
	"database/sql"
	"fmt"
	"net/http"
	"time"
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
        WHERE store_id = $1 AND deleted_at IS NULL;
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

    var existingStorePickupPlaceId int
    var deletedAt sql.NullTime
    checkPlaceQuery := `
        SELECT store_pickup_place_id, deleted_at
        FROM "store_pickup_place"
        WHERE store_id = $1 AND name = $2;
    `
    err := con.QueryRow(checkPlaceQuery, storeId, name).Scan(&existingStorePickupPlaceId, &deletedAt)
    if err != nil && err != sql.ErrNoRows{
        return res, err
    }

    if existingStorePickupPlaceId != 0 && !deletedAt.Valid {
        res.Success = false
        res.Status = http.StatusConflict
        res.Message = "Place already exist"
        return res, nil
    }

    if existingStorePickupPlaceId != 0 && deletedAt.Valid {
        updateQuery := `
            UPDATE "store_pickup_place"
            SET deleted_at = NULL
            WHERE store_pickup_place_id = $1;
        `
        stmt, err := con.Prepare(updateQuery)
        if err != nil {
            return res, err
        }
        defer stmt.Close()

        _, err = stmt.Exec(existingStorePickupPlaceId)
        if err != nil {
            return res, err
        }

        res.Success = true
        res.Status = http.StatusOK
        res.Message = "Success to restore place"
        res.Data = map[string]int{"RestoreStorePickupPlaceId": existingStorePickupPlaceId}
        return res, nil
    }

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
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add store pickup place!"
    res.Data = map[string]int64{"LastStorePickupPlaceId": id}

    return res, nil
}

func UpdateStorePickupPlace(storePickupPlaceId int, storeId int, name string) (Response, error) {
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    var existingPlaceId int
    checkPlaceQuery := `
        SELECT store_pickup_place_id
        FROM "store_pickup_place"
        WHERE store_id = $1 AND name = $2;
    `
    err := con.QueryRow(checkPlaceQuery, storeId, name).Scan(&existingPlaceId)
    if err != nil && err != sql.ErrNoRows{
        return res, err
    }

    if existingPlaceId != 0 || name == ""{
        res.Success = false
        res.Status = http.StatusBadRequest
        res.Message = "No data to update"
        return res, nil
    }

    updateStatement := `
        UPDATE "store_pickup_place"
        SET name = $2
        WHERE store_pickup_place_id = $1;
    `
    stmt, err := con.Prepare(updateStatement)
    if err != nil {
        return res, fmt.Errorf("failed to prepare update statement")
    }
    defer stmt.Close()

    result, err := stmt.Exec(storePickupPlaceId, name)
    if err != nil {
        return res, fmt.Errorf("failed to execute update statement")
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return res, fmt.Errorf("failed to retrieve affected rows")
    }
    
    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to update store pickup place"
    res.Data = map[string]int64{"rowsAffected": rowsAffected}
    
    return res, nil
}

func DeleteStorePickupPlace(storePickupPlaceId int) (Response, error) {
    var res Response
    
    con := db.CreateCon()
    // defer con.Close()
    
    deletedAt := time.Now()
    sqlStatement := `
        UPDATE "store_pickup_place"
        SET deleted_at = $2
        WHERE store_pickup_place_id = $1;
    `
    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, fmt.Errorf("failed to prepare update statement") 
    }
    
    result, err := stmt.Exec(storePickupPlaceId, deletedAt)
    if err != nil {
        return res, fmt.Errorf("failed to execute update statement")
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return res, fmt.Errorf("failed to retrieve affected rows")
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to delete store pickup place"
    res.Data = map[string]int64{"rowsAffected":rowsAffected}

    return res, nil
}   