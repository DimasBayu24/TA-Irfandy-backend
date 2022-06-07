package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string
	Price       int
	Stock       int
	Category    string
	PictureUrl  string
}
