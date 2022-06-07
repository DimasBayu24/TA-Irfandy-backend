package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Role     string
}

type UserModel struct {
	gorm.Model
	Username string
	Fullname string
}
