package entity

import (
	"coffee_api/commons"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*commons.SQLModel `json:",inline"`
	FullName          string                `json:"fullname" gorm:"column:fullname;"`
	Email             string                `json:"email" gorm:"column:email;type:varchar(100);unique_index"`
	Phone             string                `json:"phone" gorm:"column:phone;"`
	Address           string                `json:"address" gorm:"column:address;"`
	Avatar            *commons.Image        `json:"avatar" gorm:"column:avatar;"`
	Status            commons.StatusAllowed `json:"status" gorm:"column:status;type:ENUM('active','suspended','inactive');default:'active'"`
	Role              commons.RoleAllowed   `json:"role" gorm:"column:role;type:ENUM('admin','seller','shipper','member');default:'member'"`
	OTPCode           int                   `json:"otp_code" gorm:"column:otp_code"`
	IsEmailVerified   bool                  `json:"is_email_verified" gorm:"column:is_email_verified;default:false"`
	Password          string                `gorm:"column:password;" json:"-"`
	Gender            commons.Gender        `json:"gender" gorm:"column:gender;type:ENUM('male','female','other');"`
	AccessToken       string                `json:"access_token" gorm:"column:access_token;"`
	Os                *OsType               `json:"os" gorm:"column:os;"`
}

func (User) TableName() string { return "users" }

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return string(u.Role)
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) ToSimpler() *commons.SimpleUser {
	simple := &commons.SimpleUser{}
	simple.Id = user.Id
	simple.Uid = user.Uid
	// simple.Role = user.Role
	return simple
}

func (user *User) IsActive() bool {
	return user.Status == commons.Active
}
