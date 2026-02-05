package main

import (
	"coding-test/p3/config"
	"coding-test/p3/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// Routes
	r.GET("/top-spending-users", handlers.GetTopSpendingCountryUsers)
	r.POST("/users", handlers.CreateUser)

	r.Run(":8080")
}
