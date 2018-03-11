package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"gowork/controllers"
	"gowork/db"
	"gowork/helpers"
	"gowork/models"

	stdContext "context"
	"time"
)

var sessionManager *helper.SessionManager
var connection *database.Connection

func main() {
	sessionManager = helper.NewSessionManager()
	connection = database.GetInstance()
	connection.Open()
	connection.LaunchMigration(&model.User{})

	defer connection.Close()

	app := iris.New()

	iris.RegisterOnInterrupt(func() {
		fmt.Println("~~~ attemping graceful shutdown")
		defer connection.Close()
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	gate := mvc.New(app.Party("/"))
	gate.Register(sessionManager)
	gate.Handle(new(controller.MyController))

	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}
