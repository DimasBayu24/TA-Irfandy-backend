package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID   int
	ProductID int
	Quantity  int
}
