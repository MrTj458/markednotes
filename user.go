package markednotes

import (
	"time"
)

type UserService interface {
	Add(User) (User, error)
	All() ([]User, error)
	ByID(id int) (User, error)
	ByEmail(string) (User, error)
	ByUsername(string) (User, error)
	CheckInUse(User) ([]ErrorField, error)
}

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
