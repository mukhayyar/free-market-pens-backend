package models

import (
	"backend/db"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	UserId        int
	Email          string
	Username       string
	WhatsappNumber string
	Password       string
	FullName       string
	CreatedAt      string
	UpdatedAt      string
}

func GetUser(userId int) (Response, error) {
    var user User
    var res Response

    con := db.CreateCon()
    // defer con.Close()

    sqlStatement := "SELECT * FROM \"user\" where user_id = $1;"
    row := con.QueryRow(sqlStatement, userId)

    err := row.Scan(&user.UserId, &user.Email, &user.Username, &user.WhatsappNumber, &user.FullName, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return res, err
        }
        return res, err
    }

    res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to show the user data"
    res.Data = user

    return res, nil
}


func CreateUser(email string, username string, whatsappNumber string, fullName string, password string) (Response, error) {
    var res Response

    con := db.CreateCon()

    sqlStatement := "INSERT INTO \"user\" (email, username, whatsapp_number, full_name, password) VALUES($1, $2, $3, $4, $5) RETURNING user_id;"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return res, err
    }
    defer stmt.Close()

    var id int64
    err = stmt.QueryRow(email, username, whatsappNumber, fullName, password).Scan(&id)
    if err != nil {
        return res, err
    }

	res.Success = true
    res.Status = http.StatusOK
    res.Message = "Success to Add User!"
    res.Data = map[string]int64{"LastId": id}

    return res, nil
}

func UpdateUser(userId int, email string, username string, whatsappNumber string, fullName string, password string) (Response, error) {
    var res Response

    con := db.CreateCon()

    var updateValues []interface{}
    var sqlValues []string

    columns := []struct {
        name  string
        value string
    }{
        {"email", email},
        {"username", username},
        {"whatsapp_number", whatsappNumber},
        {"full_name", fullName},
        {"password", password},
    }

    for _, col := range columns {
        if col.value != "" {
            sqlValues = append(sqlValues, col.name+" = $"+strconv.Itoa(len(updateValues)+1))
            updateValues = append(updateValues, col.value)
        }
    }

    sqlStatement := "UPDATE \"user\" SET " + strings.Join(sqlValues, ", ") + " WHERE user_id = $" + strconv.Itoa(len(updateValues)+1) + ";"
    updateValues = append(updateValues, userId)

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

    res.Status = http.StatusOK
    res.Message = "Success Update User!"
    res.Data = map[string]int64{"rowsAffected   ": rowsAffected}

    return res, nil
}

func DeleteUser(userId int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM \"user\" WHERE user_id = $1;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	} 
	
	result, err := stmt.Exec(userId)
	if err != nil {
		return res, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Delete User!"
	res.Data = map[string]int64{"rows": rowsAffected}

	return res, nil
}
