package main

import (
	"restAPI/database"
	"restAPI/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()

	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders", handlers.GetOrder)
	router.PUT("/orders/:orderId", handlers.UpdateOrder)
	router.DELETE("/orders/:orderId", handlers.DeleteOrder)

	router.Run(":8000")
}