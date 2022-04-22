package markednotes

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal error")
	ErrBadRequest = errors.New("bad request")
	ErrInUse      = errors.New("in use")
)

type Error struct {
	StatusCode int
	Detail     string
	Fields     []ErrorField
}

type ErrorField struct {
	Name   string
	Type   string
	Detail string
}
