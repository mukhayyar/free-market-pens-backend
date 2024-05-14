package models

import (
	"backend/db"
	"fmt"
	"net/http"
)

type Store struct {
	StoreId         int
	UserId          int
	Name       		string
	PhotoProfile 	string
	WhatsappNumber  string
}

func GetStoreById(storeId int) (Response, error) {
    var store Store
    var products []Product
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := `
        SELECT s.store_id, s.name AS store_name, s.whatsapp_number, s.photo_profile, p.product_id, p.name AS product_name, p.price
        FROM store s
        LEFT JOIN product p ON s.store_id = p.store_id
        WHERE s.store_id = $1;
    `
    rows, err := con.Query(sqlStatement, storeId)
    if err != nil {
        return res, err
    }
    defer rows.Close()

    for rows.Next() {
        var product Product
        var productID *int
        var productName *string
        var productPrice *float64

        err := rows.Scan(&store.StoreId, &store.Name, &store.WhatsappNumber, &store.PhotoProfile, &productID, &productName, &productPrice)
        if err != nil {
            return res, err
        }
    
        if productID != nil && productName != nil && productPrice != nil{
            product.ProductId = *productID
            product.Name = *productName 
            product.Price = *productPrice 
        }
    
        products = append(products, product)
    }
    

    storeData := map[string]interface{}{
        "store_id":       store.StoreId,
        "store_name":     store.Name,
        "whatsapp_number": store.WhatsappNumber,
        "profile_photo":  store.PhotoProfile,
    }

    // productsData := map[string]interface{}{
    //     "store_id":       store.StoreId,
    //     "store_name":     store.Name,
    //     "whatsapp_number": store.WhatsappNumber,
    //     "profile_photo":  store.PhotoProfile,
    // }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = fmt.Sprintf("Profile '%s' successfully retrieved", store.Name)
    res.Data = map[string]any{
        "store": storeData,
        "products": products,
    }

    return res, nil
}

func CreateStore(userId string, Name string, PhotoProfile string, WhatsappNumber string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := "INSERT INTO \"store\" (user_id, name, photo_profile, whatsapp_number) VALUES($1, $2, $3, $4) RETURNING store_id;"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(userId, Name, PhotoProfile, WhatsappNumber).Scan(&id)
    if err != nil {
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = fmt.Sprintf("Success to create new store:%s", Name)
    res.Data = map[string]int64{"LastStoreId": id}

    return res, nil
}

// func UpdateUser(userId int, email string, username string, whatsappNumber string, fullName string, password string) (Response, error) {
//     var res Response

//     con := db.CreateCon()

//     var updateValues []interface{}
//     var sqlValues []string

//     columns := []struct {
//         name  string
//         value string
//     }{
//         {"email", email},
//         {"username", username},
//         {"whatsapp_number", whatsappNumber},
//         {"full_name", fullName},
//         {"password", password},
//     }

//     for _, col := range columns {
//         if col.value != "" {
//             sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
//             updateValues = append(updateValues, col.value)
//         }
//     }

//     sqlStatement := "UPDATE \"user\" SET " + strings.Join(sqlValues, ", ") + " WHERE user_id = $" + strconv.Itoa(len(updateValues)+1) + ";"
//     updateValues = append(updateValues, userId)

//     stmt, err := con.Prepare(sqlStatement)
//     if err != nil {
//         return res, err
//     }
//     defer stmt.Close()

//     result, err := stmt.Exec(updateValues...)
//     if err != nil {
//         return res, err
//     }

//     rowsAffected, err := result.RowsAffected()
//     if err != nil {
//         return res, err
//     }

//     res.Status = http.StatusOK
//     res.Message = "Success Update User!"
//     res.Data = map[string]int64{"rowsAffected   ": rowsAffected}

//     return res, nil
// }

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
