package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MrTj458/markednotes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Port int
	mux  *chi.Mux

	UserService   markednotes.UserService
	NoteService   markednotes.NoteService
	FolderService markednotes.FolderService
}

func NewServer(port int) *Server {
	s := &Server{
		mux:  chi.NewMux(),
		Port: port,
	}

	// Setup middleware
	s.mux.Use(middleware.Logger)
	s.mux.Use(middleware.Recoverer)

	// Add Routes
	s.addTodoRoutes()
	s.addNoteRoutes()
	s.addFoldersRoutes()

	return s
}

func (s *Server) Run() error {
	log.Println("Starting server on port:", s.Port)
	return http.ListenAndServe(":"+strconv.Itoa(s.Port), s.mux)
}
