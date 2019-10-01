package models

import (
	"github.com/clvrbarros/ForumAPI/services/users/helper"
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
func (c *Conn) CreateUser(u *User) error {
	hashPass, err := helper.HashPassword(u.Password, 10)
	if err != nil {
		return err
	}
	_, err = c.db.Exec(queryUserInsert, u.Email, hashPass, u.FirstName, u.LastName, u.Country)
	if err != nil {
		return err
	}
	return nil
}
