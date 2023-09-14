package entity

import "errors"

var (
	ErrCannotCreateUser      = errors.New("cannot create user")
	ErrEmailCannotBeBlank    = errors.New("email cannot be black")
	ErrPasswordCannotBeBlank = errors.New("password cannot be black")
	ErrPasswordLength        = errors.New("password must be at least 6 characters long")
	ErrUnauthorized          = errors.New("the account does not exist on the system")
	ErrFullNameInvalid       = errors.New("invalid string, contains other characters")
)
