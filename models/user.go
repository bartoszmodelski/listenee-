package model

import (
	"github.com/jinzhu/gorm"
	"gowork/db"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
}

func UserFirstOrCreate(email string) {
	db := database.GetInstance().DB

	var user User
	db.FirstOrCreate(&user, &User{Email: email})
}
