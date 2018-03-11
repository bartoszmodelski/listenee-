package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"gowork/controllers"
	"gowork/helpers"
)

var sessionManager *helper.SessionManager

func main() {
	sessionManager = helper.NewSessionManager()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	gate := mvc.New(app.Party("/"))
	gate.Register(sessionManager)
	gate.Handle(new(controller.MyController))

	app.Run(iris.Addr(":8080"))
}
