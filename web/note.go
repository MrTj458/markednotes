package web

import (
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

	var notes []markednotes.Note

	// Check for folder query param
	query := r.URL.Query().Get("folder")
	if len(query) > 0 {
		// Convert to int
		folderID, err := strconv.Atoi(query)
		if err != nil {
			s.renderErr(w, http.StatusBadRequest, "invalid folder ID provided")
			return
		}

		// Get the folder trying to search for
		f, err := s.FolderService.ByID(folderID)
		if err != nil {
			s.renderNotFoundOrInternal(w, err)
			return
		}

		// Check permissions
		if f.UserID != user.ID {
			s.renderErr(w, http.StatusForbidden, "you don't own that folder")
			return
		}

		// Get notes for that folder
		notes, err = s.NoteService.ByFolder(folderID)
		if err != nil {
			s.renderErrInternal(w)
			return
		}
	} else {
		// Get notes at the root folder
		var err error
		notes, err = s.NoteService.ByUserRoot(user.ID)
		if err != nil {
			s.renderErrInternal(w)
			return
		}
	}

	s.renderJSON(w, http.StatusOK, notes)
}

func (s *Server) handleNotesCreate(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	type NoteIn struct {
		FolderID *int   `json:"folder_id"`
		Name     string `json:"name" validate:"required,min=1,max=30"`
		Body     string `json:"body"`
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

	// Check folder permissions
	if note.FolderID != nil {
		f, err := s.FolderService.ByID(*note.FolderID)
		if err != nil {
			s.renderNotFoundOrInternal(w, err)
			return
		}

		if f.UserID != user.ID {
			s.renderErr(w, http.StatusForbidden, "you do not own that folder")
			return
		}
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
		s.renderNotFoundOrInternal(w, err)
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
		s.renderNotFoundOrInternal(w, err)
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
		s.renderNotFoundOrInternal(w, err)
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
