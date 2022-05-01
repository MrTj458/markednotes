package markednotes

import "time"

type Folder struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ParentID  *int      `json:"parent_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FolderService interface {
	Add(*Folder) error
	All() ([]Folder, error)
	ByID(int) (Folder, error)
	ByUser(int) ([]Folder, error)
	ByParent(int) ([]Folder, error)
	ByUserRoot(int) ([]Folder, error)
	Update(*Folder) error
	Delete(int) error
}
