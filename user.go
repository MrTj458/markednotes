package markednotes

import "time"

type User struct {
	ID        uint64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface{}
