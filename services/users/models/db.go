package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Database *sql.DB
var err error

func init() {
	Database, err = sql.Open("postgres", "user=postgres dbname=forum password=admin host=127.0.0.1 sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	if err = Database.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("Database connected")
}