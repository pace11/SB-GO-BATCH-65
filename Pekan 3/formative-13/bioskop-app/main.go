package main

import (
	"bioskop-app/database"
	"bioskop-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()

	router.POST("/bioskop", handlers.CreateBioskop)

	router.Run(":8080")
}
