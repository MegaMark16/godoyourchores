package handlers

import (
	"fmt"
	"html/template"
	"net/http"
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
