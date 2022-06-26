package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     int
	OrderDate  time.Time
	Status     string
	TotalPrice int
	PaymentUrl string
	OrderItem  []OrderItem
}
