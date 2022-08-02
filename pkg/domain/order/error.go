package order

import "github.com/pkg/errors"

var (
	ErrEmptyRepository = errors.New("repository cannot be nil")
)
