package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) addFoldersRoutes() {
	s.mux.Route("/api/folders", func(r chi.Router) {
		r.Get("/", s.handleFoldersIndex)
	})
}

func (s *Server) handleFoldersIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Folders index"))
}
