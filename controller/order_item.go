package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
)

func GetAllOrderItems(c *gin.Context) {
	var product []model.OrderItem
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetOrderItemByID(c *gin.Context) {
	var product model.OrderItem
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetOrderItemByOrderID(c *gin.Context) {
	type Result struct {
		ID          int    `gorm:"column:id"`
		OrderID     int    `gorm:"column:order_id"`
		Quantity    int    `gorm:"column:quantity"`
		ProductName string `gorm:"column:product_name"`
	}
	var result []Result
	err := db.DB.Raw("SELECT a.id, a.order_id, a.quantity, b.product_name FROM order_items a "+
		"LEFT JOIN products b "+
		"ON a.product_id = b.id "+
		"LEFT JOIN orders c "+
		"ON a.order_id = c.id "+
		"WHERE c.id = ?", c.Query("order_id")).Scan(&result).Error
	if err != nil {
		fmt.Println("error")

	}
	fmt.Println("tes hasil order user id", c.Query("user_id"))
	fmt.Println("tes hasil order get", result)
	c.JSON(http.StatusOK, result)
}

func PostOrderItem(c *gin.Context) {
	var input form.OrderItem
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.OrderItem{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateOrderItemByID(c *gin.Context) {
	var product model.OrderItem
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.OrderItem
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteOrderItemByID(c *gin.Context) {
	var product model.OrderItem
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
