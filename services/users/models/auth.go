package models

import (
	"errors"
	"github.com/clvrbarros/ForumAPI/services/users/helper"
	"github.com/dgrijalva/jwt-go"
)

// Auth type is a struct for authentication
type Auth struct {
	Email 	 string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthClaims type is a struct for JWT Claims
type AuthClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// AuthJWT type is a struct for JWT authentication
type AuthJWT struct {
	Token string `json:"token"`
}

func (c *Conn) Authenticate(a *Auth) (bool, error) {
	var dbPassword string

	row := c.db.QueryRow(queryAuthPassword, a.Email)
	err := row.Scan(&dbPassword)
	if err != nil {
		return false, errors.New("E-mail não cadastrado")
	}

	checkPassword := helper.CheckPasswordHash(a.Password, dbPassword)
	if !checkPassword {
		return false, nil
	}

	return true, nil
}
