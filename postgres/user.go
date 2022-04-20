package postgres

import (
	"github.com/MrTj458/markednotes"
	"github.com/jackc/pgx/v4"
)

var _ markednotes.UserService = (*UserService)(nil)

type UserService struct {
	db *pgx.Conn
}

func NewUserService(db *pgx.Conn) *UserService {
	return &UserService{
		db: db,
	}
}
