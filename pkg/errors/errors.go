package errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrBadID    = errors.New("bad id")
	ErrFormat   = errors.New("not supported format")
)
