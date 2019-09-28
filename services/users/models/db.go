package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var err error

func InitDb() error {
	db, err = sql.Open("postgres", "user=postgres dbname=forum password=admin host=127.0.0.1 sslmode=disable")
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("Database connected")
	return nil
}

func CloseDb() {
	db.Close()
}