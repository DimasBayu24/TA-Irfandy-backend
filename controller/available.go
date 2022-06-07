package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GetAllAvailables(c *gin.Context) {
	var product []model.Available
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetAvailableByID(c *gin.Context) {
	var product model.Available
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProductByDay(c *gin.Context) {
	type Result struct {
		ID          int    `gorm:"column:id"`
		ProductName string `gorm:"column:product_name"`
		Price       int    `gorm:"column:price"`
		Stock       int    `gorm:"column:stock"`
		PictureUrl  string `gorm:"column:picture_url"`
	}
	var result []Result
	//var product model.Available
	db.DB.Table("products").Select("products.id, products.product_name, products.price, products.stock, products.picture_url").Joins("left join availables on products.id = availables.product_id").Where("availables.day = ?", c.Query("day")).Where("availables.deleted_at IS NULL").Scan(&result)
	fmt.Println(result)

	c.JSON(http.StatusOK, result)
}

func GetAvailableByProductID(c *gin.Context) {
	var product []model.Available
	if err := db.DB.Where("product_id = ?", c.Query("product_id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostAvailable(c *gin.Context) {
	var input form.Available
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("var1 = ", reflect.ValueOf(input.Day).Kind())
	fmt.Println(input.Day)
	fmt.Println(json.Marshal(input.Day))
	for _, element := range input.Day {
		product := model.Available{
			ProductID: input.ProductID,
			Day:       string(element),
		}
		db.DB.Create(&product)
	}

	c.JSON(http.StatusCreated, "OK")
}

func UpdateAvailableByID(c *gin.Context) {
	var product model.Available
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.Available
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteAvailableByID(c *gin.Context) {
	var product model.Available
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
