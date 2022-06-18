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

func GetAllOrderByID(c *gin.Context) {
	var product []model.Order
	//var productItem model.OrderItem
	//type ResultItem struct {
	//	ID int `gorm:"column:id"`
	//	ProductName string `gorm:"column:product_name"`
	//}
	//
	//type Result struct {
	//	ID int `gorm:"column:id"`
	//	ResultItem []ResultItem
	//	Status string `gorm:"column:status"`
	//	TotalPrice int `gorm:"column:total_price"`
	//	OrderDate string `gorm:"column:order_date"`
	//}
	//
	//var resultItem []ResultItem
	//var result Result
	//
	//db.DB.Table("order_items").Select("order_items.id, products.product_name").Joins("left join products on products.id = order_items.product_id").Where("order_items.order_id = ?", c.Query("order_id")).Scan(&resultItem)
	//db.DB.Table("orders").Select("orders.id, orders.status, orders.total_price, orders.order_date").Where("user_id = ?",c.Query("user_id")).Scan(&result)
	db.DB.Preload("OrderItem").Find(&product).Where("user_id = ?", c.Query("user_id"))

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
