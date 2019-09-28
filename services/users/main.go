package main

import (
	"github.com/clvrbarros/ForumAPI/services/users/controllers"
	"github.com/clvrbarros/ForumAPI/services/users/models"
	"github.com/clvrbarros/ForumAPI/services/users/router"
	"log"
)

func main() {
	controllers.InitValidator()    // Init validator instance
	err := models.InitDb()		   // Init db instance
	if err != nil {
		log.Panic(err)
	}
	log.Println("PostgresSQL: Connection OK")
	defer models.CloseDb()

	r := router.NewRouter()			// Load routes
	r.Run()
}