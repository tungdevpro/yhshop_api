package entity

import "errors"


var (
	ErrCannotCreateUser = errors.New("Cannot create user")
	ErrEmailCannotBeBlank = errors.New("Email cannot be black")
	ErrPasswordCannotBeBlank = errors.New("Password cannot be black")
)


