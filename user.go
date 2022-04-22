package markednotes

import (
	"regexp"
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

type UserIn struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserIn) ToUser() User {
	return User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u UserIn) Validate() ([]ErrorField, bool) {
	var errors []ErrorField

	if len(u.Username) < 3 || len(u.Username) > 15 {
		err := ErrorField{
			Name:   "username",
			Type:   "string",
			Detail: "Username must be between 3 and 15 characters.",
		}
		errors = append(errors, err)
	}

	emailRegex := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	if !emailRegex.Match([]byte(u.Email)) {
		err := ErrorField{
			Name:   "email",
			Type:   "string",
			Detail: "Email must be a valid email address.",
		}
		errors = append(errors, err)
	}

	if len(u.Password) < 6 {
		err := ErrorField{
			Name:   "password",
			Type:   "string",
			Detail: "Password must be at least 6 characters long.",
		}
		errors = append(errors, err)
	}

	if len(errors) == 0 {
		return nil, true
	}
	return errors, false
}
