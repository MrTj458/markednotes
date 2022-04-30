package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrTj458/markednotes"
	"github.com/go-chi/chi/v5"
)

func (s *Server) addFoldersRoutes() {
	s.mux.Route("/api/folders", func(r chi.Router) {
		r.Use(s.requireAuth)
		r.Get("/", s.handleFoldersIndex)
		r.Post("/", s.handleFoldersCreate)
		r.Get("/{id}", s.handleFolderById)
		r.Put("/{id}", s.handleFolderUpdate)
		r.Delete("/{id}", s.handleFoldersDelete)
	})
}

func (s *Server) handleFoldersIndex(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	folders, err := s.FolderService.ByUser(user.ID)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, folders)
}

func (s *Server) handleFoldersCreate(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(markednotes.User{}).(markednotes.User)

	type FolderIn struct {
		ParentID *int   `json:"parent_id"`
		Name     string `json:"name" validate:"required,min=2,max=30"`
	}

	// Decode JSON
	var folderIn FolderIn
	if err := s.decodeJSON(w, r.Body, &folderIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(folderIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid folder received", errors)
		return
	}

	folder := markednotes.Folder{
		UserID:   user.ID,
		ParentID: folderIn.ParentID,
		Name:     folderIn.Name,
	}

	// Add folder to DB
	err := s.FolderService.Add(&folder)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, folder)
}

func (s *Server) handleFolderById(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	folder, err := s.FolderService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("folder with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if folder.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that folder")
		return
	}

	s.renderJSON(w, http.StatusOK, folder)
}

func (s *Server) handleFolderUpdate(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	type FolderIn struct {
		ParentID *int   `json:"parent_id"`
		Name     string `json:"name" validate:"required,min=2,max=30"`
	}

	// Decode JSON
	var folderIn FolderIn
	if err := s.decodeJSON(w, r.Body, &folderIn); err != nil {
		s.renderErr(w, http.StatusBadRequest, "invalid JSON received")
		return
	}

	// Validate
	if errors, ok := s.Validator.Struct(folderIn); !ok {
		s.renderErrFields(w, http.StatusBadRequest, "invalid folder received", errors)
		return
	}

	// Get folder
	folder, err := s.FolderService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("folder with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if folder.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that folder")
		return
	}

	folder.ParentID = folderIn.ParentID
	folder.Name = folderIn.Name

	err = s.FolderService.Update(&folder)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	s.renderJSON(w, http.StatusOK, folder)
}

func (s *Server) handleFoldersDelete(w http.ResponseWriter, r *http.Request) {
	user := s.getUserFromRequest(r)

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		s.renderErr(w, http.StatusBadRequest, "ID must be an integer")
		return
	}

	folder, err := s.FolderService.ByID(id)
	if err != nil {
		switch err {
		case markednotes.ErrNotFound:
			s.renderErr(w, http.StatusNotFound, fmt.Sprintf("folder with ID '%d' not found", id))
		default:
			s.renderErrInternal(w)
		}
		return
	}

	// Check permissions
	if folder.UserID != user.ID {
		s.renderErr(w, http.StatusForbidden, "You do not own that folder")
		return
	}

	err = s.FolderService.Delete(id)
	if err != nil {
		s.renderErrInternal(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
