package controllers

import (
	"database/sql"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	h = initControllers()
	router = gin.Default()
)

func initControllers() *Setup {
	databaseInstance := databaseConn()
	modelInstance := models.Connect(databaseInstance)
	validatorInstance := validator.New()
	controllerInstance := NewController(modelInstance, validatorInstance)

	return controllerInstance
}

func databaseConn() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=forum " +
		"password=admin host=127.0.0.1 sslmode=disable")
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

func request(method string, url string, body io.Reader, t *testing.T) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}

	return req, httptest.NewRecorder()
}