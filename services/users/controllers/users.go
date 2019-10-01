package controllers

import (
	"github.com/clvrbarros/ForumAPI/services/users/helper"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *Setup) CreateUser (c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, &helper.APIMessage{Message:err.Error()})
	}

	if err := s.validator.Struct(request); err != nil {
		helper.ValidatorMessage(c, err)
		return
	}

	err := s.model.CreateUser(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &helper.APIMessage{Message:"Internal Server Error"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, &helper.APIMessage{Message: "User created successfully"})
}