package form

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     int       `json:"UserID" binding:"required"`
	OrderDate  time.Time `json:"OrderDate" binding:"required"`
	Status     string    `json:"Status" binding:"required"`
	TotalPrice int       `json:"TotalPrice" binding:"required"`
}
