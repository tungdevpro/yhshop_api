package entity

import "coffee_api/commons"

type Category struct {
	*commons.SQLModel `json:",inline"`
	Name              string         `json:"name" gorm:"column:name;"`
	Description       string         `json:"description" gorm:"column:description;"`
	Image             *commons.Image `json:"avatar" gorm:"column:avatar;"`
}
