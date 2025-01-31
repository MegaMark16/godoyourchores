package handlers

import (
	"log"
	"net/http"

	"github.com/megamark16/godoyourchores/internal/models"
	"github.com/megamark16/godoyourchores/pkg/auth"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", PageData{})
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
