package web

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/MrTj458/markednotes"
)

func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from header
		token := r.Header.Get("Authorization")
		if len(token) == 0 || !strings.Contains(token, " ") {
			s.renderErr(w, http.StatusUnauthorized, "invalid token provided")
			return
		}

		// Get ID from token
		id, err := s.Jwt.Parse(strings.Split(token, " ")[1])
		if err != nil {
			switch err {
			case markednotes.ErrTokenExpired:
				s.renderErr(w, http.StatusUnauthorized, "token expired")
				return
			default:
				s.renderErr(w, http.StatusUnauthorized, "invalid token provided")
				return
			}
		}

		// Find user with ID in token
		user, err := s.UserService.ByID(id)
		if err != nil {
			log.Println("requireAuth:", err)
			s.renderErrInternal(w)
			return
		}

		// Pass user on to next handler
		ctx := context.WithValue(r.Context(), markednotes.User{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
