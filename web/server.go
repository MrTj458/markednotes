package web

import (
	"encoding/json"
	"io"
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

	Validator markednotes.Validator

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

func (s *Server) decodeJSON(w http.ResponseWriter, r io.Reader, out any) {
	dec := json.NewDecoder(r)
	if err := dec.Decode(out); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
	}
}

func (s *Server) renderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.Encode(data)
}

func (s *Server) renderErr(w http.ResponseWriter, status int, detail string) {
	res := markednotes.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     make([]markednotes.ErrorField, 0),
	}
	s.renderJSON(w, status, res)
}

func (s *Server) renderErrFields(w http.ResponseWriter, status int, detail string, fields []markednotes.ErrorField) {
	res := markednotes.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     fields,
	}
	s.renderJSON(w, status, res)
}

func (s *Server) renderErrInternal(w http.ResponseWriter) {
	s.renderErr(w, http.StatusInternalServerError, "internal server error")
}
