package controllers

import (
	"errors"
	"github.com/clvrbarros/ForumAPI/services/users/helper"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Setup) AuthUser(c *gin.Context) {
	var request models.Auth
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, &helper.APIMessage{Message:err.Error()})
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(c, err)
		return
	}

	checkPassword, err := s.model.Authenticate(&request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &helper.APIMessage{Message:err.Error()})
		return
	}
	if !checkPassword {
		c.JSON(http.StatusUnauthorized, &helper.APIMessage{Message:"Senha incorreta!"})
		return
	}

	token, err := s.GenerateToken(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &helper.APIMessage{Message:err.Error()})
	}

	userInfo := models.AuthJWT{Token: token}

	c.JSON(http.StatusOK, &userInfo)
}
func (s *Setup) GenerateToken(email string) (string, error) {
	secret := []byte("SECRET#@")

	var claims models.AuthClaims

	claims.Email = email
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 360).Unix(),
		Issuer: "ForumAPI",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("Erro na geração do token")
	}

	return signedString, nil
}
