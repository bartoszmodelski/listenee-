package main

import (
	"./models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Connection struct {
	db *DB
}

func (connection *Connection) Open() {
	var err error
	connection.db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	connection.db.AutoMigrate(&model.User{})
}

func (connection *Connection) Close() {
	connection.db.Close()
}
