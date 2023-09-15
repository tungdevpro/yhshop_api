package entity

import "coffee_api/commons"

type Shop struct {
	*commons.SQLModel `json:",inline"`
	Name              string         `json:"name" gorm:"column:name"`
	OwnerId           int            `json:"owner_id" gorm:"column:owner_id;"`
	CityId            int            `json:"city_id" gorm:"column:city_id;"`
	Icon              *commons.Image `json:"icon" gorm:"column:icon"`
	Image             *commons.Image `json:"image" gorm:"column:image"`
	Status            int            `json:"status" gorm:"column:status;default:'1'"`
}

func (s *Shop) Mask(isOwner bool) {
	s.GenUID()
}
