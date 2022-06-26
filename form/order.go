package form

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     int       `json:"UserID"`
	OrderDate  time.Time `json:"OrderDate"`
	Status     string    `json:"Status"`
	TotalPrice int       `json:"TotalPrice"`
	PaymentUrl string    `json:"PaymentUrl"`
}
