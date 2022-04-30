package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrTj458/markednotes"
	"github.com/go-chi/chi/v5"
)

func (s *Server) addNoteRoutes() {
	s.mux.Route("/api/notes", func(r chi.Router) {
		r.Use(s.requireAuth)
		r.Get("/", s.handleNotesIndex)
		r.Post("/", s.handleNotesCreate)
		r.Get("/{id}", s.handleNoteById)
		r.Put("/{id}", s.handleNoteUpdate)
		r.Delete("/{id}", s.handleNoteDelete)
	})
}

func (s *Server) handleNotesIndex(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	notes, err := s.NoteService.ByUser(user.ID)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, notes)
}

func (s *Server) handleNotesCreate(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	type NoteIn struct {
		FolderID *int   `json:"folder_id"`
		Name     string `json:"name" validate:"required,min=2,max=30"`
		Body     string `json:"body" validate:"required"`
	}

	// Decode JSON
	var noteIn NoteIn
	if err := s.decodeJSON(w, r.Body, &noteIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(noteIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid note received", errors)
		return
	}

	note := markednotes.Note{
		UserID:   user.ID,
		FolderID: noteIn.FolderID,
		Name:     noteIn.Name,
		Body:     noteIn.Body,
	}

	// Add note to DB
	err := s.NoteService.Add(&note)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, note)
}

func (s *Server) handleNoteById(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	note, err := s.NoteService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("note with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if note.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that note")
		return
	}

	s.renderJSON(w, http.StatusOK, note)
}

func (s *Server) handleNoteUpdate(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	type NoteIn struct {
		FolderID *int   `json:"folder_id"`
		Name     string `json:"name" validate:"required,min=2,max=30"`
		Body     string `json:"body" validate:"required"`
	}

	// Decode JSON
	var noteIn NoteIn
	if err := s.decodeJSON(w, r.Body, &noteIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(noteIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid folder received", errors)
		return
	}

	// Get note
	note, err := s.NoteService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("note with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if note.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that note")
		return
	}

	note.FolderID = noteIn.FolderID
	note.Name = noteIn.Name
	note.Body = noteIn.Body

	err = s.NoteService.Update(&note)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, note)
}

func (s *Server) handleNoteDelete(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	note, err := s.NoteService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("note with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if note.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that note")
		return
	}

	err = s.NoteService.Delete(id)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
