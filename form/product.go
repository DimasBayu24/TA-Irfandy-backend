package form

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"ProductName" binding:"required"`
	Price       int    `json:"Price" binding:"required"`
	Stock       int    `json:"Stock" binding:"required"`
	Category    string `json:"Category" binding:"required"`
	PictureUrl  string `json:"PictureUrl" binding:"required"`
}
