package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    "sync"
)

// Connection body

type Connection struct {
	DB *gorm.DB
}

func (connection *Connection) Open() {
	var err error
	connection.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
}

func (connection *Connection) Close() {
	connection.DB.Close()
}

func (connection *Connection) LaunchMigration(models ...interface{}) {
	connection.DB.AutoMigrate(models...)
}

// Singleton

var instance *Connection
var once sync.Once

func GetInstance() *Connection {
    once.Do(func() {
        instance = &Connection{}
    })
    return instance
}