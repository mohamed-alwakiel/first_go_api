package main

import (
	"example/firstApp/controllers"
	"example/firstApp/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize router
	r := gin.Default()

	// connect to data base
	database.ConnectDatabase()

	// routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// listen and serve on port [8080]
	r.Run()
}
