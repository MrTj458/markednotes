package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) addNoteRoutes() {
	s.mux.Route("/api/notes", func(r chi.Router) {
		r.Get("/", s.handleNotesIndex)
	})
}

func (s *Server) handleNotesIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Notes index"))
}
