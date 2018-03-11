package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Cookie string
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
