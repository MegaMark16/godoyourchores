package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/megamark16/godoyourchores/internal/handlers"
	"github.com/megamark16/godoyourchores/pkg/auth"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/chores", handlers.ChoresHandler).Methods("GET")
	r.HandleFunc("/complete-chore", handlers.CompleteChoreHandler).Methods("POST")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	// Initialize authentication
	auth.InitializeSessionStore()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
