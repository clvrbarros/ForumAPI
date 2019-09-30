package main

import (
	"database/sql"
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/clvrbarros/ForumAPI/services/users/router"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

func main() {
	databaseInstance := database("user=postgres dbname=forum " +
		"password=admin host=127.0.0.1 sslmode=disable")
	modelInstance := models.Connect(databaseInstance)
	validatorInstance := validator.New()
	controllerInstance := controllers.NewController(modelInstance, validatorInstance)

	defer databaseInstance.Close()

	r := router.NewRouter(controllerInstance)
	r.Run()
}

func database(connect string) *sql.DB {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatalln("Could not open connection to Postgres: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Could not connect to MySQL: ", err)
	}

	log.Println("Database connected")
	return db
}