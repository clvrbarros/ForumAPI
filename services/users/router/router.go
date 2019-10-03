package router

import (
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(c *controllers.Setup) *gin.Engine {
	router := gin.Default()

	// Routes
	router.POST("/register", c.CreateUser)
	router.POST("/login", c.AuthUser)

	return router
}
