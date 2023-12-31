package entity

import (
	"coffee_api/commons"
)

type ListShopEmpty []Shop

type Shop struct {
	*commons.SQLModel `json:",inline"`
	Name              string              `json:"name" gorm:"column:name;"`
	UserId            int                 `json:"-" gorm:"column:owner_id;"`
	CityId            int                 `json:"city_id" gorm:"column:city_id;"`
	Icon              *commons.Image      `json:"icon" gorm:"column:icon"`
	Image             *commons.Image      `json:"images" gorm:"column:images"`
	Status            int                 `json:"status" gorm:"column:status;default:1;"`
	IsVerify          int                 `json:"is_verify" gorm:"column:is_verify;default:1;"`
	LikedCount        int                 `json:"liked_count" gorm:"column:liked_count;"`
	User              *commons.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Shop) TableName() string {
	return "shops"
}
func (s *Shop) Mask(isOwner bool) {
	s.GenerateID()

	if u := s.User; u != nil {
		u.GenerateID()
	}
}

type Filter struct {
	IsVerify int `json:"is_verify,omitempty" form:"is_verify"`
	Status   int `json:"status,omitempty" form:"status"`
}
