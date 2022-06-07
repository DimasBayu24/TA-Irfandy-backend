package form

import (
	"gorm.io/gorm"
)

type Available struct {
	gorm.Model
	ProductID int   `json:"ProductID" binding:"required"`
	Day      []Days `json:"Day" binding:"required"`
}

type Days string
