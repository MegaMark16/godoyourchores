package handlers

import (
	"net/http"

	"github.com/megamark16/godoyourchores/pkg/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if auth.Authenticate(username, password) {
			auth.SetSession(username, w, r)
			http.Redirect(w, r, "/chores", http.StatusSeeOther)
			return
		}
		renderTemplate(w, "login", PageData{Message: "Invalid credentials"})
		return
	}
	renderTemplate(w, "login", PageData{})
}
