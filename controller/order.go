package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
)

func GetAllOrders(c *gin.Context) {
	var product []model.Order
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostOrder(c *gin.Context) {
	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Order{
		UserID:     input.UserID,
		OrderDate:  input.OrderDate,
		Status:     input.Status,
		TotalPrice: input.TotalPrice,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Order
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteOrderByID(c *gin.Context) {
	var product model.Order
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
