package postgres

import (
	"github.com/MrTj458/markednotes"
	"github.com/jackc/pgx/v4"
)

var _ markednotes.FolderService = (*FolderService)(nil)

type FolderService struct {
	db *pgx.Conn
}

func NewFolderService(db *pgx.Conn) *FolderService {
	return &FolderService{
		db: db,
	}
}
