package main

import (
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
)

func main() {
	defer models.Database.Close()
	r := gin.Default()
	r.POST("/register", controllers.CreateUser)
	r.Run()
}