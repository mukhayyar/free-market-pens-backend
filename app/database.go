package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=Mantablah1 dbname=free-market sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_{nama-tabel}
	// migrate -database "postgres://postgres:Mantablah1@localhost:5432/free-market-test?sslmode=disable" -path db/migrations up
}