package entity

import (
	"coffee_api/commons"
	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"
)

type StatusAllowed string

const (
	Suspended StatusAllowed = "suspended"
	Active    StatusAllowed = "active"
	Inactive  StatusAllowed = "inactive"
)

func (r *StatusAllowed) Scan(value interface{}) error {
	*r = StatusAllowed(value.([]byte))
	return nil
}

func (r StatusAllowed) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	*commons.SQLModel `json:",inline"`
	FullName          string         `json:"fullname" gorm:"column:fullname;"`
	Email             string         `json:"email" gorm:"column:email;type:varchar(100);unique_index"`
	Phone             string         `json:"phone" gorm:"column:phone;"`
	Address           string         `json:"address" gorm:"column:address;"`
	Avatar            *commons.Image `json:"avatar" gorm:"column:avatar;"`
	Status            StatusAllowed  `json:"status" gorm:"column:status;type:ENUM('active','suspended','inactive');default:'active'"`
	Role              RoleAllowed    `json:"role" gorm:"column:role;type:ENUM('admin','seller','shipper','member');default:'member'"`
	OTPCode           int            `json:"otp_code" gorm:"column:otp_code"`
	IsEmailVerified   bool           `json:"is_email_verified" gorm:"column:is_email_verified;default:false"`
	Password          []byte         `gorm:"not null" json:"-"`

	AccessToken string `json:"access_token" gorm:"column:access_token;"`
}

func (u *User) TableName() string {
	return "users"
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
