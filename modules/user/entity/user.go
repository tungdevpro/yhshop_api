package entity

import "coffee_api/commons"

type User struct {
	*commons.SQLModel `json:",inline"`
	FullName          string `json:"fullname"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	Avatar            string `json:"avatar"`
	Status            string `json:"status"`
}
