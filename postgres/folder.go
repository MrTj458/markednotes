package postgres

import (
	"context"
	"log"
	"time"

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

func (fs *FolderService) Add(folder *markednotes.Folder) error {
	sql := `
		INSERT INTO folders (parent_id, user_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)
		RETURNING id, created_at, updated_at
	`

	now := time.Now().UTC()

	row := fs.db.QueryRow(context.Background(), sql, folder.ParentID, folder.UserID, folder.Name, now)

	err := row.Scan(&folder.ID, &folder.CreatedAt, &folder.UpdatedAt)
	if err != nil {
		log.Println("FolderService.Add:", err)
		return markednotes.ErrInternal
	}

	return nil
}

func (fs *FolderService) All() ([]markednotes.Folder, error) {
	sql := `
		SELECT * FROM folders
	`

	rows, err := fs.db.Query(context.Background(), sql)
	if err != nil {
		log.Println("FolderService.All:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	folders := make([]markednotes.Folder, 0)
	for rows.Next() {
		var f markednotes.Folder
		err = rows.Scan(
			&f.ID,
			&f.ParentID,
			&f.UserID,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
		)
		if err != nil {
			log.Println("FolderService.All:", err)
			return nil, markednotes.ErrInternal
		}
		folders = append(folders, f)
	}

	return folders, nil
}

func (fs *FolderService) ByID(id int) (markednotes.Folder, error) {
	sql := `
		SELECT * FROM folders
		WHERE id = $1
	`

	row := fs.db.QueryRow(context.Background(), sql, id)

	var f markednotes.Folder
	err := row.Scan(
		&f.ID,
		&f.ParentID,
		&f.UserID,
		&f.Name,
		&f.CreatedAt,
		&f.UpdatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return markednotes.Folder{}, markednotes.ErrNotFound
		default:
			log.Println("UserService.ByID:", err)
			return markednotes.Folder{}, markednotes.ErrInternal
		}
	}

	return f, nil
}

func (fs *FolderService) ByUser(userId int) ([]markednotes.Folder, error) {
	sql := `
		SELECT * FROM folders
		WHERE user_id = $1
	`

	rows, err := fs.db.Query(context.Background(), sql, userId)
	if err != nil {
		log.Println("FolderService.ByUser:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	folders := make([]markednotes.Folder, 0)
	for rows.Next() {
		var f markednotes.Folder
		err = rows.Scan(
			&f.ID,
			&f.ParentID,
			&f.UserID,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
		)
		if err != nil {
			log.Println("FolderService.ByUser:", err)
			return nil, markednotes.ErrInternal
		}
		folders = append(folders, f)
	}

	return folders, nil
}

func (fs *FolderService) Update(folder *markednotes.Folder) error {
	sql := `
		UPDATE folders
		SET name = $2, parent_id = $3, updated_at = $4
		WHERE id = $1
		RETURNING updated_at
	`

	now := time.Now().UTC()

	row := fs.db.QueryRow(context.Background(), sql, folder.ID, folder.Name, folder.ParentID, now)
	err := row.Scan(&folder.UpdatedAt)
	if err != nil {
		log.Println("FolderService.Update:", err)
		return markednotes.ErrInternal
	}

	return nil
}

func (fs *FolderService) Delete(id int) error {
	sql := `
		DELETE from folders
		WHERE id = $1
	`

	_, err := fs.db.Exec(context.Background(), sql, id)
	if err != nil {
		log.Println("FolderService.Delete:", err)
		return markednotes.ErrInternal
	}
	return nil
}
