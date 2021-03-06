package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MrTj458/markednotes"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) addUserRoutes() {
	s.mux.Route("/api/users", func(r chi.Router) {
		r.Post("/", s.handleUsersCreate)
		r.Get("/", s.handleUsersIndex)
		r.Get("/{id}", s.handleUserByID)
		r.Post("/login", s.handleUserLogin)

		// Authenticated routes
		r.Group(func(a chi.Router) {
			a.Use(s.requireAuth)
			a.Get("/me", s.handleUserMe)
		})
	})
}

func (s *Server) handleUsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserService.All()
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, users)
}

func (s *Server) handleUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	user, err := s.UserService.ByID(id)
	if err != nil {
		s.renderNotFoundOrInternal(w, err)
		return
	}

	s.renderJSON(w, http.StatusOK, user)
}

func (s *Server) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
	type UserIn struct {
		Username string `json:"username" validate:"min=2,max=15"`
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"min=6"`
	}

	// Decode JSON
	var userIn UserIn
	if err := s.decodeJSON(w, r.Body, &userIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(userIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid user received", errors)
		return
	}

	user := markednotes.User{
		Username: userIn.Username,
		Email:    userIn.Email,
		Password: userIn.Password,
	}

	// Check if email or username is in use
	if errors, err := s.UserService.CheckInUse(user); err != nil {
		switch err {
		case markednotes.ErrInUse:
			s.renderErrFields(w, http.StatusBadRequest, "invalid user received", errors)
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Bcrypt Hash:", err)
		s.renderErrInternal(w)
		return
	}
	user.Password = string(hashed)

	// Add user to DB
	err = s.UserService.Add(&user)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	// Generate JWT
	token, err := s.Jwt.NewToken(user.ID)
	if err != nil {
		log.Println("JWT:", err)
		s.renderErrInternal(w)
		return
	}

	res := map[string]any{
		"token": token,
		"user":  user,
	}

	s.renderJSON(w, http.StatusCreated, res)
}

func (s *Server) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	type UserIn struct {
		Email    string `json:"email" validate:"email"`
		Password string `json:"password" validate:"min=6"`
	}

	// Decode JSON
	var userIn UserIn
	if err := s.decodeJSON(w, r.Body, &userIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(userIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid user received", errors)
		return
	}

	// Find user with given email
	user, err := s.UserService.ByEmail(userIn.Email)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusUnauthorized, "invalid email or password")
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userIn.Password))
	if err != nil {
		s.renderErr(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Generate JWT
	token, err := s.Jwt.NewToken(user.ID)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	res := map[string]any{
		"token": token,
		"user":  user,
	}

	s.renderJSON(w, http.StatusOK, res)
}

func (s *Server) handleUserMe(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)
	s.renderJSON(w, http.StatusOK, user)
}
