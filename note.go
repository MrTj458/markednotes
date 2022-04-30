package markednotes

import "time"

type Note struct {
	ID        int
	UserID    int
	FolderID  *int
	Name      string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NoteService interface {
	Add(*Note) error
	All() ([]Note, error)
	ByID(int) (Note, error)
	ByUser(int) ([]Note, error)
	ByFolder(int) ([]Note, error)
	Update(*Note) error
	Delete(int) error
}
