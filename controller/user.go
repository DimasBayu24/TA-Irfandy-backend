package controller

import (
	"net/http"
	"product-api/db"
	"product-api/form"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Phone:    input.Phone,
		Role:     input.Role,
	}
	db.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func GetUserByUsername(c *gin.Context) {
	var user model.User
	if err := db.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func LoginAdmin(c *gin.Context) {
	var user model.User
	if err := db.DB.Where("username = ?", c.Query("username")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Record": "Record not found"})
		return
	}

	if user.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		return
	}
}

func GetAllUsers(c *gin.Context) {
	var product []model.User
	db.DB.Find(&product)

	c.JSON(http.StatusOK, product)
}

func GetUserByID(c *gin.Context) {
	var product model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func PostUser(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := model.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Phone:    input.Phone,
		Role:     input.Role,
	}
	db.DB.Create(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateUserByID(c *gin.Context) {
	var product model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, product)
}

func DeleteUserByID(c *gin.Context) {
	var product model.User
	if err := db.DB.Where("id = ?", c.Query("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
