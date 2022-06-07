package form

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   int `json:"OrderID" binding:"required"`
	ProductID int `json:"ProductID" binding:"required"`
	Quantity  int `json:"Quantity" binding:"required"`
}
