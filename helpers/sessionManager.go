package helper

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	"time"
)

type SessionManager struct {
	session *sessions.Sessions
}

func NewSessionManager() *SessionManager {
	sessionDB, _ := boltdb.New("./sessions.db", 0666, "users")
	// use different go routines to sync the database
	sessionDB.Async(true)

	// close and unlock the database when control+C/cmd+C pressed
	iris.RegisterOnInterrupt(func() {
		sessionDB.Close()
	})

	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	sessionManager := new(SessionManager)
	sessionManager.session = sessions.New(sessions.Config{
		Cookie:  "session",
		Expires: -1 * time.Minute, // <=0 means unlimited life
		Encode:  secureCookie.Encode,
		Decode:  secureCookie.Decode,
	})

	sessionManager.session.UseDatabase(sessionDB)

	return sessionManager
}

func (sessionManager *SessionManager) setVariable(ctx context.Context, key string, value string) {
	s := sessionManager.session.Start(ctx)
	s.Set(key, value)
}

func (sessionManager *SessionManager) getVariable(ctx context.Context, key string) string {
	return sessionManager.session.Start(ctx).GetString(key)
}

func (sessionManager *SessionManager) removeVariable(ctx context.Context, key string) {
	sessionManager.session.Start(ctx).Delete("name")
}

func (sessionManager *SessionManager) renew(ctx context.Context) {
	sessionManager.session.ShiftExpiration(ctx)
}

func (sessionManager *SessionManager) end(ctx context.Context) {
	sessionManager.session.Destroy(ctx)
}

func (sessionManager *SessionManager) removeAllVariables(ctx context.Context) {
	sessionManager.session.Start(ctx).Clear()
}

func (sessionManager *SessionManager) LogIn(ctx context.Context, email string) {
	sessionManager.setVariable(ctx, "email", email)
}

func (sessionManager *SessionManager) IsLoggedIn(ctx context.Context) bool {
	return sessionManager.getVariable(ctx, "email") != ""
}

func (sessionManager *SessionManager) GetEmail(ctx context.Context) string {
	return sessionManager.getVariable(ctx, "email")
}

func (sessionManager *SessionManager) LogOut(ctx context.Context) {
	sessionManager.end(ctx)
}