package entity

import (
	"net/mail"
	"regexp"
	"strings"
)

type RegisterDTO struct {
	FullName        string `json:"fullname" form:"fullname"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
	Token           string `json:"token,omitempty"`
	Id              string `json:"id,omitempty"`
}

func (r *RegisterDTO) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if err := isEmailAddress(r.Email); err != nil {
		return err
	}

	r.Password = strings.TrimSpace(r.Password)
	if len(r.Password) < 6 {
		return ErrPasswordLength
	}

	r.FullName = strings.TrimSpace(r.FullName)
	if len(r.FullName) != 0 {
		regex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
		if err := regex.MatchString(r.FullName); !err {
			return ErrFullNameInvalid
		}
	}
	return nil
}

type LoginDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (r *LoginDTO) Validate() error {
	r.Email = strings.TrimSpace(r.Email)
	if err := isEmailAddress(r.Email); err != nil {
		return err
	}

	r.Password = strings.TrimSpace(r.Password)
	if len(r.Password) < 6 {
		return ErrPasswordLength
	}

	return nil
}

func isEmailAddress(input string) error {
	_, err := mail.ParseAddress(input)

	return err
}

type CreateUser struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Email       string `json:"email" gorm:"column:email;"`
	FullName    string `json:"fullname" gorm:"column:fullname;"`
	Password    string `json:"-" gorm:"column:password;"`
	AccessToken string `json:"access_token"`
}

func (CreateUser) TableName() string { return "users" }
