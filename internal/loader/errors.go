package loader

import "errors"

var (
	ErrNoGoMod      = errors.New("go.mod not found")
	ErrEmptyPattern = errors.New("empty package pattern")
)