package constant

import "errors"

var (
	// ErrUserNotFound is a constant of error message when user not found
	ErrUserNotFound = errors.New("user not found")
	ErrPassword     = errors.New("password not match")
	ErrNotPremium   = errors.New("user not premium")
)
