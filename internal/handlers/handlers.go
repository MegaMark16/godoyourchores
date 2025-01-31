package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/megamark16/godoyourchores/internal/models"
	"github.com/megamark16/godoyourchores/pkg/auth"
)

type PageData struct {
	Username string
	Chores   []string
	Message  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	t.Execute(w, data)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", PageData{})
}

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

func ChoresHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := auth.GetSession(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "chores", PageData{Username: username, Chores: models.GetChores(username)})
}

func CompleteChoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		chore := r.Form.Get("chore")
		day := r.Form.Get("day")
		log.Printf("Chore '%s' completed for day '%s'", chore, day)
		http.Redirect(w, r, "/chores", http.StatusSeeOther)
	}
}
