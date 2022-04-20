package postgres

import (
	"github.com/MrTj458/markednotes"
	"github.com/jackc/pgx/v4"
)

var _ markednotes.UserService = (*NoteService)(nil)

type NoteService struct {
	db *pgx.Conn
}

func NewNoteService(db *pgx.Conn) *NoteService {
	return &NoteService{
		db: db,
	}
}
