package markednotes

import "time"

type Folder struct {
	ID        uint64
	UserID    uint64
	ParentID  uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FolderService interface{}
