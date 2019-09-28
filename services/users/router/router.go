package router

import (
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Routes
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.CreateUser)
	router.POST("/register2", controllers.CreateUser)

	return router
}
