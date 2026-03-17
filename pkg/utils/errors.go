package utils

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrDatabaseError = errors.New("database error")
)