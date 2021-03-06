package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MrTj458/markednotes"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ markednotes.UserService = (*UserService)(nil)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		db: db,
	}
}

// Add creates a new row in the database sing the values from the given User.
// It adds all of the generated fields to the passed in User.
func (us *UserService) Add(user *markednotes.User) error {
	sql := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)
		RETURNING id, created_at, updated_at
	`

	now := time.Now().UTC()

	row := us.db.QueryRow(context.Background(), sql, user.Username, user.Email, user.Password, now)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("UserService.Add:", err)
		return markednotes.ErrInternal
	}

	return nil
}

// All returns all of the User's that are stored in the database.
func (us *UserService) All() ([]markednotes.User, error) {
	sql := `
		SELECT * FROM users
	`

	rows, err := us.db.Query(context.Background(), sql)
	if err != nil {
		log.Println("UserService.All:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	users := make([]markednotes.User, 0)
	for rows.Next() {
		var u markednotes.User
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			log.Println("UserService.All:", err)
			return nil, markednotes.ErrInternal
		}
		users = append(users, u)
	}

	return users, nil
}

// ByID returns a User with the given ID.
func (us *UserService) ByID(id int) (markednotes.User, error) {
	sql := `
		SELECT * FROM users
		WHERE id = $1
	`

	row := us.db.QueryRow(context.Background(), sql, id)

	var u markednotes.User
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return markednotes.User{}, markednotes.ErrNotFound
		default:
			log.Println("UserService.ByID:", err)
			return markednotes.User{}, markednotes.ErrInternal
		}
	}

	return u, nil
}

// ByEmail returns a User with the given email.
func (us *UserService) ByEmail(email string) (markednotes.User, error) {
	sql := `
		SELECT * FROM users
		WHERE email = $1
	`

	row := us.db.QueryRow(context.Background(), sql, email)

	var u markednotes.User
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return markednotes.User{}, markednotes.ErrNotFound
		default:
			log.Println("UserService.ByEmail:", err)
			return markednotes.User{}, markednotes.ErrInternal
		}
	}

	return u, nil
}

// ByUsername returns a User with the given username.
func (us *UserService) ByUsername(username string) (markednotes.User, error) {
	sql := `
		SELECT * FROM users
		WHERE username = $1
	`

	row := us.db.QueryRow(context.Background(), sql, username)

	var u markednotes.User
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return markednotes.User{}, markednotes.ErrNotFound
		default:
			log.Println("UserService.ByUsername:", err)
			return markednotes.User{}, markednotes.ErrInternal
		}
	}

	return u, nil
}

// CheckInUse checks the email and username of the passed in user to make sure
// neither of them are already in use.
func (us *UserService) CheckInUse(user markednotes.User) ([]markednotes.ErrorField, error) {
	sql := `
		SELECT * FROM users
		WHERE username = $1 or email = $2
	`

	rows, err := us.db.Query(context.Background(), sql, user.Username, user.Email)
	if err != nil {
		log.Println("UserService.CheckTakenAccount:", err)
		return nil, markednotes.ErrInternal
	}
	defer rows.Close()

	var errors []markednotes.ErrorField
	for rows.Next() {
		var u markednotes.User
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			log.Println("UserService.CheckTakenAccount:", err)
			return nil, markednotes.ErrInternal
		}

		if user.Username == u.Username {
			newErr := markednotes.ErrorField{
				Name:   "username",
				Detail: fmt.Sprintf("username '%s' is already in use", user.Username),
			}
			errors = append(errors, newErr)
		}

		if user.Email == u.Email {
			newErr := markednotes.ErrorField{
				Name:   "email",
				Detail: fmt.Sprintf("email '%s' is already in use", user.Email),
			}
			errors = append(errors, newErr)
		}
	}

	if len(errors) == 0 {
		return nil, nil
	}
	return errors, markednotes.ErrInUse
}
