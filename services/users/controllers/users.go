package controllers

import (
	"github.com/clvrbarros/ForumAPI/services/users/helper"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser (c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(request); err != nil {
		helper.ValidatorMessage(c, err)
		return
	}

	c.JSON(http.StatusCreated, &helper.APIMessage{Message: "User created successfully"})
}