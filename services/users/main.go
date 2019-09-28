package main

import (
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	controllers.InitValidator()
	err := models.InitDb()
	if err != nil {
		log.Panic(err)
	}
	log.Println("PostgresSQL: Connection OK")
	defer models.CloseDb()

	r := gin.Default()

	// Routes
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.CreateUser)
	r.POST("/register2", controllers.CreateUser)

	r.Run()
}