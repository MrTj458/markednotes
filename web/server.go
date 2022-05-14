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
	Jwt       markednotes.Jwt

	UserService   markednotes.UserService
	NoteService   markednotes.NoteService
	FolderService markednotes.FolderService
}

// NewServer creates a new Server instance with preconfigured values.
func NewServer(port int) *Server {
	s := &Server{
		mux:  chi.NewMux(),
		Port: port,
	}

	// Setup middleware
	s.mux.Use(middleware.Logger)
	s.mux.Use(middleware.Recoverer)

	// Add Routes
	s.addUserRoutes()
	s.addNoteRoutes()
	s.addFoldersRoutes()

	// Static files
	s.mux.Handle("/*", http.FileServer(http.Dir("static")))

	return s
}

// Run starts the server
func (s *Server) Run() error {
	log.Println("Starting server on port:", s.Port)
	return http.ListenAndServe(":"+strconv.Itoa(s.Port), s.mux)
}

// decodeJSON will decode the value in r into the given struct out.
func (s *Server) decodeJSON(w http.ResponseWriter, r io.Reader, out any) error {
	dec := json.NewDecoder(r)
	if err := dec.Decode(out); err != nil {
		return err
	}
	return nil
}

// renderJSON takes in a status code and a struct and encodes it into the given
// http.ResponseWriter.
func (s *Server) renderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.Encode(data)
}

// renderErr will send a JSON formatted error response.
func (s *Server) renderErr(w http.ResponseWriter, status int, detail string) {
	res := markednotes.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     make([]markednotes.ErrorField, 0),
	}
	s.renderJSON(w, status, res)
}

// renderErrFields will send a JSON formatted error response, along with the
// given slice of error fields.
func (s *Server) renderErrFields(w http.ResponseWriter, status int, detail string, fields []markednotes.ErrorField) {
	res := markednotes.Error{
		StatusCode: status,
		Detail:     detail,
		Fields:     fields,
	}
	s.renderJSON(w, status, res)
}

// renderErrInternal is a shortcut to return a 500 internal server error.
func (s *Server) renderErrInternal(w http.ResponseWriter) {
	s.renderErr(w, http.StatusInternalServerError, "internal server error")
}

func (s *Server) renderNotFoundOrInternal(w http.ResponseWriter, e error) {
	switch e {
	case markednotes.ErrNotFound:
		s.renderErr(w, http.StatusNotFound, "not found")
	default:
		s.renderErrInternal(w)
	}
}

func (s *Server) getUserFromRequest(r *http.Request) markednotes.User {
	return r.Context().Value(markednotes.User{}).(markednotes.User)
}
