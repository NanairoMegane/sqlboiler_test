package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	dsn := "test_u:test_pw@tcp(localhost:3306)/boiler_test"

	openedDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	openedDB.SetMaxIdleConns(10)
	openedDB.SetMaxOpenConns(10)
	openedDB.SetConnMaxLifetime(300 * time.Second)

	DB = openedDB
}
