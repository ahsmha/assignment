package api

import (
	"net/http"

	"github.com/ahsmha/assignment/api/handlers"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", handlers.SignupHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/notes", handlers.NotesHandler)
	mux.HandleFunc("/notes", handlers.CreateNoteHandler).Methods("POST")
	mux.HandleFunc("/notes", handlers.DeleteNoteHandler).Methods("DELETE")

	return mux
}
