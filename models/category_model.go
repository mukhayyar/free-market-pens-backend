package models

import (
	"backend/db"
	"net/http"
)

type Category struct {
	CategoryId        int
	Name          string
}

func GetAllCategory() (Response, error) {
	var category Category
	var categoryList []Category
	var res Response

	con := db.CreateCon()
	// defer con.Close()

	sqlStatement := "SELECT * FROM \"category\";"
	rows, err := con.Query(sqlStatement)
	if err != nil{
		return res, err
	}
	defer rows.Close()
	
	for rows.Next(){
		err = rows.Scan(&category.CategoryId, &category.Name)
		if err != nil{
			return res, err
		}
		categoryList = append(categoryList, category)
	}

    res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show the list of category"
	res.Data = categoryList

	return res, nil
}


func CreateCategory(name string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := "INSERT INTO \"category\" (name) VALUES($1) RETURNING category_id;"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(name).Scan(&id)
    if err != nil {
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to add new category"
    res.Data = map[string]int64{"LastId": id}

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
