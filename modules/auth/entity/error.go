package entity

import "errors"

var (
	ErrCannotCreateUser      = errors.New("cannot create user")
	ErrEmailCannotBeBlank    = errors.New("email cannot be black")
	ErrPasswordCannotBeBlank = errors.New("password cannot be black")
	ErrPasswordLength        = errors.New("password must be at least 6 characters long")
	ErrUnauthorized          = errors.New("the account does not exist on the system")
	ErrFullNameInvalid       = errors.New("invalid string, contains other characters")
	ErrVerifiedYourAccount   = errors.New("you have not verified your account")
	ErrOTPLength             = errors.New("otp must be 5 characters long")
	ErrOTPNotEqual           = errors.New("OTP does not match")
)

type AuthCode int

var (
	Success     AuthCode = 1
	Error       AuthCode = -1
	NotVerified AuthCode = -4
)
