package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root:["PASSWORD"]@tcp(127.0.0.1:3306)/["DATABASE_NAME"]")
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	return db
}
