package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var sessionManager *SessionManager

func main() {
	sessionManager = NewSessionManager()
	// the rest of the code stays the same.
	app := iris.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})
	app.Get("/set", func(ctx context.Context) {
		sessionManager.SetVariable(ctx, "name", "default")
		ctx.Writef("name = default")
	})

	app.Get("/set/{key}/{value}", func(ctx context.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		sessionManager.SetVariable(ctx, key, value)
		// test if setted here
		ctx.Writef("%s = %s", key, value)
	})

	app.Get("/get", func(ctx context.Context) {
		name := sessionManager.GetVariable(ctx, "name")

		ctx.Writef("The name on the /set was: %s", name)
	})

	app.Get("/get/{key}", func(ctx context.Context) {
		key := ctx.Params().Get("key")
		value := sessionManager.GetVariable(ctx, key)
		ctx.Writef("%s = %s", key, value)
	})

	app.Get("/delete/{key}", func(ctx context.Context) {
		key := ctx.Params().Get("key")
		sessionManager.RemoveVariable(ctx, key)
		ctx.Writef("Deleted key: %s", key)
	})

	app.Get("/clear", func(ctx context.Context) {
		sessionManager.RemoveAllVariables(ctx)
		ctx.Writef("Removed all variables")
	})

	app.Get("/destroy", func(ctx context.Context) {
		sessionManager.End(ctx)
		ctx.Writef("Session ended")
	})

	app.Get("/update", func(ctx context.Context) {
		sessionManager.Renew(ctx)
		ctx.Writef("Session renewed")
	})

	app.Run(iris.Addr(":8080"))
}
