package entity

import (
	"net/mail"
	"strings"
)

type RegisterDTO struct {
	FullName string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token,omitempty"`
	Id       string `json:"id,omitempty"`
}

type LoginDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (r *RegisterDTO) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if err := isEmailAddress(r.Email); err != nil {
		return err
	}

	r.Password = strings.TrimSpace(r.Password)
	if len(r.Password) < 6 {
		return ErrorPasswordLength
	}

	return nil
}

func isEmailAddress(input string) error {
	_, err := mail.ParseAddress(input)

	return err
}
