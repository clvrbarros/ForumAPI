package models

import (
	"log"
	"time"
)

// User type is a struct for users table
type User struct {
	ID 		  int 		`json:"id"`
	Email     string 	`json:"email" validate:"required,email"`
	Password  string 	`json:"password" validate:"required,min=6"`
	FirstName string 	`json:"firstname" validate:"required"`
	LastName  string 	`json:"lastname"`
	Country   string 	`json:"country"`
	Active 	  bool 		`json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateUser creates a new user
func CreateUser() error {
	sqlStmt := "INSERT INTO test (name) values ('clevinho')"
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
