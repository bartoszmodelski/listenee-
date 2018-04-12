package controller

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"gowork/helpers"
	"gowork/models"
)

type GateController struct {
	Session *helper.SessionManager
}

func (m *GateController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login/{email:string}", "GetLogin")
	b.Handle("GET", "/secret", "GetSecret", m.AuthenticationGate)
}

// GET: http://localhost:8080/
func (m *GateController) Get(ctx context.Context) string {
	if m.Session.IsLoggedIn(ctx) {
		ctx.ViewData("Title", "Hi Page")
		ctx.ViewData("Name", m.Session.GetEmail(ctx))
		ctx.View("hi.html")
		return ""
	} else {
		return "please login"
	}
}

func (m *GateController) GetLogin(ctx context.Context, email string) string {
	m.Session.LogIn(ctx, email)
	model.UserFirstOrCreate(email)
	return m.Get(ctx)
}

func (m *GateController) AuthenticationGate(ctx context.Context) {
	if m.Session.IsLoggedIn(ctx) {
		ctx.Next()
	} else {
		ctx.Redirect("/login")
	}
}

func (m *GateController) GetLogout(ctx context.Context) string {
	m.Session.LogOut(ctx)
	return "logged out"
}

func (m *GateController) GetSecret(ctx context.Context) string {
	return "some secret"
}
