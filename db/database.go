package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init() {
	db, err = sql.Open("postgres", "user=postgres password=123 dbname=free-market sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxIdleTime(10 * time.Minute)
	// db.SetConnMaxLifetime(60 * time.Minute)

	// migrate create -ext sql -dir db/migrations create_table_{nama-tabel}
	// migrate -database "postgres://postgres:Mantablah1@localhost:5432/free-market-test?sslmode=disable" -path db/migrations up
}

func CreateCon() *sql.DB {
	return db
}
