package postgres

import (
	"context"
	"log"
	"time"

	"github.com/MrTj458/markednotes"
	"github.com/jackc/pgx/v4"
)

var _ markednotes.NoteService = (*NoteService)(nil)

type NoteService struct {
	db *pgx.Conn
}

func NewNoteService(db *pgx.Conn) *NoteService {
	return &NoteService{
		db: db,
	}
}

func (ns *NoteService) Add(note *markednotes.Note) error {
	sql := `
	INSERT INTO notes (user_id, folder_id, name, body, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $5)
	RETURNING id, created_at, updated_at
`

	now := time.Now().UTC()

	row := ns.db.QueryRow(context.Background(), sql, note.UserID, note.FolderID, note.Name, note.Body, now)

	err := row.Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		log.Println("NoteService.Add:", err)
		return markednotes.ErrInternal
	}

	return nil
}

func (ns *NoteService) All() ([]markednotes.Note, error) {
	sql := `
		SELECT * FROM notes
	`

	rows, err := ns.db.Query(context.Background(), sql)
	if err != nil {
		log.Println("NoteService.All:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	notes := make([]markednotes.Note, 0)
	for rows.Next() {
		var n markednotes.Note
		err = rows.Scan(
			&n.ID,
			&n.UserID,
			&n.FolderID,
			&n.Name,
			&n.Body,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		if err != nil {
			log.Println("NoteService.All:", err)
			return nil, markednotes.ErrInternal
		}
		notes = append(notes, n)
	}

	return notes, nil
}

func (ns *NoteService) ByID(id int) (markednotes.Note, error) {
	sql := `
		SELECT * FROM notes
		WHERE id = $1
	`

	row := ns.db.QueryRow(context.Background(), sql, id)

	var n markednotes.Note
	err := row.Scan(
		&n.ID,
		&n.UserID,
		&n.FolderID,
		&n.Name,
		&n.Body,
		&n.CreatedAt,
		&n.UpdatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return markednotes.Note{}, markednotes.ErrNotFound
		default:
			log.Println("NoteService.ByID:", err)
			return markednotes.Note{}, markednotes.ErrInternal
		}
	}

	return n, nil
}

func (ns *NoteService) ByUser(userId int) ([]markednotes.Note, error) {
	sql := `
		SELECT * FROM notes
		WHERE user_id = $1
	`

	rows, err := ns.db.Query(context.Background(), sql, userId)
	if err != nil {
		log.Println("NoteService.ByUser:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	notes := make([]markednotes.Note, 0)
	for rows.Next() {
		var n markednotes.Note
		err = rows.Scan(
			&n.ID,
			&n.UserID,
			&n.FolderID,
			&n.Name,
			&n.Body,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		if err != nil {
			log.Println("NoteService.ByUser:", err)
			return nil, markednotes.ErrInternal
		}
		notes = append(notes, n)
	}

	return notes, nil
}

func (ns *NoteService) ByUserRoot(userId int) ([]markednotes.Note, error) {
	sql := `
		SELECT * FROM notes
		WHERE user_id = $1 and folder_id IS NULL
	`

	rows, err := ns.db.Query(context.Background(), sql, userId)
	if err != nil {
		log.Println("NoteService.ByUser:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	notes := make([]markednotes.Note, 0)
	for rows.Next() {
		var n markednotes.Note
		err = rows.Scan(
			&n.ID,
			&n.UserID,
			&n.FolderID,
			&n.Name,
			&n.Body,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		if err != nil {
			log.Println("NoteService.ByUser:", err)
			return nil, markednotes.ErrInternal
		}
		notes = append(notes, n)
	}

	return notes, nil
}

func (ns *NoteService) ByFolder(folderId int) ([]markednotes.Note, error) {
	sql := `
		SELECT * FROM notes
		WHERE folder_id = $1
	`

	rows, err := ns.db.Query(context.Background(), sql, folderId)
	if err != nil {
		log.Println("NoteService.ByFolder:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	notes := make([]markednotes.Note, 0)
	for rows.Next() {
		var n markednotes.Note
		err = rows.Scan(
			&n.ID,
			&n.UserID,
			&n.FolderID,
			&n.Name,
			&n.Body,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		if err != nil {
			log.Println("NoteService.ByFolder:", err)
			return nil, markednotes.ErrInternal
		}
		notes = append(notes, n)
	}

	return notes, nil
}

func (ns *NoteService) Update(note *markednotes.Note) error {
	sql := `
		UPDATE notes
		SET name = $2, folder_id = $3, body = $4, updated_at = $5
		WHERE id = $1
		RETURNING updated_at
	`

	now := time.Now().UTC()

	row := ns.db.QueryRow(context.Background(), sql, note.ID, note.Name, note.FolderID, note.Body, now)
	err := row.Scan(&note.UpdatedAt)
	if err != nil {
		log.Println("NoteService.Update:", err)
		return markednotes.ErrInternal
	}

	return nil
}

func (ns *NoteService) Delete(id int) error {
	sql := `
		DELETE from notes
		WHERE id = $1
	`

	_, err := ns.db.Exec(context.Background(), sql, id)
	if err != nil {
		log.Println("NoteService.Delete:", err)
		return markednotes.ErrInternal
	}
	return nil
}
