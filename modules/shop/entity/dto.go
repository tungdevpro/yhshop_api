package entity

import "coffee_api/commons"

type CreateShopDTO struct {
	Name    string         `json:"name" gorm:"column:name"`
	OwnerId string         `json:"owner_id" gorm:"column:owner_id"`
	Icon    *commons.Image `json:"icon" gorm:"column:icon"`
	CityId  int            `json:"city_id" gorm:"column:city_id"`
}
