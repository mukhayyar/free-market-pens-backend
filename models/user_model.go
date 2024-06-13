package models

import (
	"backend/db"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId         int
	Email          string
	Username       string
	WhatsappNumber string
	Password       string
	FullName       string
	CreatedAt      string
	UpdatedAt      string
	IsAdmin        bool
	StoreId        sql.NullInt64
}

func GetUser(userId int) (Response, error) {
	var user User
	var res Response
	var storeId sql.NullInt64

	con := db.CreateCon()
	// defer con.Close()

	// Modified SQL statement to also fetch the store_id from the store table
	sqlStatement := `
		SELECT u.user_id, u.username, u.email, u.whatsapp_number, u.is_admin, u.password, u.created_at, u.updated_at, s.store_id
		FROM "user" u
		LEFT JOIN "store" s ON u.user_id = s.user_id
		WHERE u.user_id = $1;
	`
	row := con.QueryRow(sqlStatement, userId)

	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.WhatsappNumber, &user.IsAdmin, &user.Password, &user.CreatedAt, &user.UpdatedAt, &storeId)
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

func CreateUser(email string, username string, whatsappNumber string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	fmt.Println("Email:", email)
	fmt.Println("Username:", username)
	fmt.Println("WhatsappNumber:", whatsappNumber)
	fmt.Println("Password:", password)

	sqlStatement := `
        INSERT INTO "user" (email, username, whatsapp_number, password) 
        VALUES($1, $2, $3, $4) 
        RETURNING user_id;
    `

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(email, username, whatsappNumber, password).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Constraint {
			case "unique_email":
				return res, fmt.Errorf("a user with the same email already exists")
			case "unique_username":
				return res, fmt.Errorf("a user with the same username already exists")
			}
		}
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to Add User!"
	res.Data = map[string]string{"email": email, "username": fmt.Sprint(id)}

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

	sqlStatement := `UPDATE "user" SET ` + strings.Join(sqlValues, ", ") + ` WHERE "user_id" = $` + strconv.Itoa(len(updateValues)+1) + ";"
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

	sqlStatement := `DELETE FROM "user" WHERE "user_id" = $1;`

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

func GetUserByEmail(email string) (Response, error) {
	var user User
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM "user" WHERE "email" = $1;`
	row := con.QueryRow(sqlStatement, email)

	err := row.Scan(&user.UserId, &user.Email, &user.Username, &user.WhatsappNumber, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin)
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

func GetUserByUsername(username string) (Response, error) {
	var user User
	var res Response

	con := db.CreateCon()
	sqlStatement := `SELECT * FROM "user" where "username" = $1;`
	row := con.QueryRow(sqlStatement, username)

	err := row.Scan(&user.UserId, &user.Email, &user.Username, &user.WhatsappNumber, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin)
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
func GetAllUsers() (Response, error) {
	var users []User
	var res Response

	con := db.CreateCon()
	sqlStatement := `
        SELECT u.user_id, u.email, u.username, u.whatsapp_number, u.created_at, u.updated_at, u.is_admin, s.store_id
        FROM "user" u
        LEFT JOIN "store" s ON u.user_id = s.user_id;
    `
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Email, &user.Username, &user.WhatsappNumber, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.StoreId)
		if err != nil {
			return res, err
		}
		users = append(users, user)
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to show all users"
	res.Data = users

	return res, nil
}

func UpdatePassword(userId int, newPassword string) (Response, error) {
	var res Response

	con := db.CreateCon()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}

	sqlStatement := `UPDATE "user" SET password = $1 WHERE user_id = $2;`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(string(hashedPassword), userId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update password"
	res.Data = map[string]int64{"rowsAffected": rowsAffected}

	return res, nil
}

func UpdateIsAdmin(userId int, isAdmin bool) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `UPDATE "user" SET is_admin = $1 WHERE user_id = $2;`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(isAdmin, userId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update is_admin status"
	res.Data = map[string]int64{"rowsAffected": rowsAffected}

	return res, nil
}

func UpdateWhatsappNumber(userId int, newWhatsappNumber string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := `UPDATE "user" SET whatsapp_number = $1 WHERE user_id = $2;`
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newWhatsappNumber, userId)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Success = true
	res.Status = http.StatusOK
	res.Message = "Success to update whatsapp number"
	res.Data = map[string]int64{"rowsAffected": rowsAffected}

	return res, nil
}
