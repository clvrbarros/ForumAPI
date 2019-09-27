package controllers

import (
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func CreateUser (c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var customError string
			switch err.Tag() {
			case "required":
				customError = "is required"
			case "email":
				customError = "is not valid"
			case "min":
				customError = "minimum length is " + err.Param()
			}
			log.Println("The " + strings.ToLower(err.Field()) + " field " + customError)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}