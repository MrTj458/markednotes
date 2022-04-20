package markednotes

import "time"

type Note struct {
	ID        uint64
	UserID    uint64
	FolderID  uint64
	Name      string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NoteService interface{}
