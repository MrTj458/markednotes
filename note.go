package markednotes

import "time"

type Note struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FolderID  *int      `json:"folder_id"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteService interface {
	Add(*Note) error
	All() ([]Note, error)
	ByID(int) (Note, error)
	ByUser(int) ([]Note, error)
	ByFolder(int) ([]Note, error)
	ByUserRoot(int) ([]Note, error)
	Update(*Note) error
	Delete(int) error
}
