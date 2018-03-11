package main

import (
	"./helpers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

var sessionManager *helper.SessionManager

func main() {
	sessionManager = helper.NewSessionManager()
	app := iris.New()
	mvc.Configure(app.Party("/"), myMVC)
	app.Run(iris.Addr(":8080"))
}

func myMVC(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know
	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	//b.Router().Use(m.authenticationGate)
	b.Handle("GET", "/login/{email:string}", "GetLogin")
	b.Handle("GET", "/secret", "GetSecret", m.authenticationGate)
}

// GET: http://localhost:8080/
func (m *MyController) Get(ctx context.Context) string {
	if sessionManager.IsLoggedIn(ctx) {
		return "logged in, email: " + sessionManager.GetEmail(ctx)
	} else {
		return "please login"
	}
}

func (m *MyController) GetLogin(ctx context.Context, email string) string {
	sessionManager.LogIn(ctx, email)
	return m.Get(ctx)
}

func (m *MyController) authenticationGate(ctx context.Context) {
	if sessionManager.IsLoggedIn(ctx) {
		ctx.Next()
	} else {
		ctx.Redirect("/login")
	}
}

func (m *MyController) GetLogout(ctx context.Context) string {
	sessionManager.LogOut(ctx)
	return "logged out"
}

func (m *MyController) GetSecret(ctx context.Context) string {
	return "some secret"
}
