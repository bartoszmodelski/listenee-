package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func findOrCreate(name string) *User {
	db.First(&product, "code = ?", "L1212") // find product with code l1212
}
