package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) addTodoRoutes() {
	s.mux.Route("/api/users", func(r chi.Router) {
		r.Get("/", s.handleUsersIndex)
	})
}

func (s *Server) handleUsersIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users index"))
}
