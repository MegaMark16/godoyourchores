package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

var users = map[string]string{
	"user": "test",
}

func InitializeSessionStore() {

}

func initSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session")
	if session.IsNew {
		session.Options.Domain = "localhost"
		session.Options.MaxAge = 10
		session.Options.HttpOnly = true
		session.Options.Secure = false
		fmt.Println("Create New Session (cookie)")
	} else {
		fmt.Println("Use Old Session (old cookie)")
	}
	return session
}

func Authenticate(username, password string) bool {
	if pass, ok := users[username]; ok && pass == password {
		return true
	}
	return false
}

func SetSession(username string, w http.ResponseWriter, r *http.Request) {
	session := initSession(r)
	session.Values["username"] = username
	session.Save(r, w)
}

func GetSession(r *http.Request) (string, bool) {
	session := initSession(r)
	username, ok := session.Values["username"].(string)
	return username, ok
}
