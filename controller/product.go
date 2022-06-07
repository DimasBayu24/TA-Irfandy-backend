package controller

import (
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var product []model.Product
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetProductByID(c *gin.Context) {
	var product model.Product
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func PostProduct(c *gin.Context) {
	var input form.Product
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Product{
		ProductName: input.ProductName,
		Price:       input.Price,
		Stock:       input.Stock,
		Category:    input.Category,
		PictureUrl:  input.PictureUrl,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateProductByID(c *gin.Context) {
	var product model.Product
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Product
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteProductByID(c *gin.Context) {
	var product model.Product
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
