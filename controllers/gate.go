package controller

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"gowork/helpers"
)

type MyController struct {
	Session *helper.SessionManager
}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login/{email:string}", "GetLogin")
	b.Handle("GET", "/secret", "GetSecret", m.AuthenticationGate)
}

// GET: http://localhost:8080/
func (m *MyController) Get(ctx context.Context) string {
	if m.Session.IsLoggedIn(ctx) {
		return "logged in, email: " + m.Session.GetEmail(ctx)
	} else {
		return "please login"
	}
}

func (m *MyController) GetLogin(ctx context.Context, email string) string {
	m.Session.LogIn(ctx, email)
	return m.Get(ctx)
}

func (m *MyController) AuthenticationGate(ctx context.Context) {
	if m.Session.IsLoggedIn(ctx) {
		ctx.Next()
	} else {
		ctx.Redirect("/login")
	}
}

func (m *MyController) GetLogout(ctx context.Context) string {
	m.Session.LogOut(ctx)
	return "logged out"
}

func (m *MyController) GetSecret(ctx context.Context) string {
	return "some secret"
}
