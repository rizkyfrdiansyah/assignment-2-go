package handlers

import (
	"net/http"
	"restAPI/database"
	"restAPI/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	orderId := c.Param("orderId")
	if err := database.DB.First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	if err := c.ShouldBindJSON(&order);  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	var order models.Order
	if err := database.DB.First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	database.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}