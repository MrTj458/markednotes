package markednotes

import "errors"

var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")
)

type Jwt interface {
	NewToken(int) (string, error)
	Parse(string) (int, error)
}
