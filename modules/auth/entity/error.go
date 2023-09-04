package entity

import "errors"

var (
	ErrCannotCreateUser      = errors.New("Cannot create user")
	ErrEmailCannotBeBlank    = errors.New("Email cannot be black")
	ErrPasswordCannotBeBlank = errors.New("Password cannot be black")
	ErrorPasswordLength      = errors.New("Password must be at least 6 characters long")
)
