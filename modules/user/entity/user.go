package entity

import (
	"coffee_api/commons"
	"database/sql/driver"
)

type UserStatus string

const (
	ACTIVE  UserStatus = "active"
	BAN     UserStatus = "ban"
	DELETED UserStatus = "deleted"
)

func (s *UserStatus) Scan(value interface{}) error {
	*s = UserStatus(value.([]byte))
	return nil
}

func (s UserStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type User struct {
	*commons.SQLModel `json:",inline"`
	FullName          string         `json:"fullname" gorm:"column:fullname;"`
	Email             string         `json:"email" gorm:"column:email;type:varchar(100);unique_index"`
	Phone             string         `json:"phone" gorm:"column:phone;"`
	Address           string         `json:"address" gorm:"column:address;"`
	Avatar            *commons.Image `json:"avatar" gorm:"column:avatar;"`
	Status            UserStatus     `json:"status" gorm:"column:status;"`
	AccessToken       string         `json:"access_token" gorm:"column:access_token;"`
}

func (u *User) TableName() string {
	return "users"
}
