package markednotes

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal error")
	ErrBadRequest = errors.New("bad request")
	ErrInUse      = errors.New("in use")
)

type Error struct {
	StatusCode int          `json:"status_code"`
	Detail     string       `json:"detail"`
	Fields     []ErrorField `json:"fields"`
}

type ErrorField struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
