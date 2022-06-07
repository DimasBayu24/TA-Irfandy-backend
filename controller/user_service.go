package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-api/db"
	"product-api/model"
	"reflect"
	"time"
)

func PostOrderFood(c *gin.Context) {
	var order model.Order
	type Input struct {
		ID       int `json:"id" binding:"required"`
		UserID   int `json:"user_id" binding:"required"`
		Quantity int `json:"quantity" binding:"required"`
		Price    int `json:"price" binding:"required"`
	}
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("var1 = ", reflect.ValueOf(input.Quantity).Kind())
	fmt.Println(input.Quantity)
	fmt.Println(json.Marshal(input))

	db.DB.First(&order).Where("user_id = ?",input.UserID)
	if order.ID == 0 {
		fmt.Println(order.ID == 0)
		product := model.Order{
			UserID:     input.UserID,
			OrderDate:  time.Now(),
			Status:     "Not Done",
			TotalPrice: input.Price,
		}
		db.DB.Create(&product)

		orderItem := model.OrderItem{
			OrderID:   int(product.ID),
			ProductID: input.ID,
			Quantity:  input.Quantity,
		}

		db.DB.Create(&orderItem)
	} else {
		orderItem := model.OrderItem{
			OrderID:   int(order.ID),
			ProductID: input.ID,
			Quantity:  input.Quantity,
		}

		db.DB.Create(&orderItem)

		orderUpdate := model.Order{
			UserID:     order.UserID,
			OrderDate:  order.OrderDate,
			Status:     "Not Done",
			TotalPrice: order.TotalPrice + input.Price,
		}
		db.DB.Model(&order).Updates(orderUpdate)

	}

	fmt.Println(order.ID)
	fmt.Println(order.UserID)

	c.JSON(http.StatusCreated, order)
}
