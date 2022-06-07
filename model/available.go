package model

import (
	"gorm.io/gorm"
)

type Available struct {
	gorm.Model
	ProductID int
	Day       string
}
