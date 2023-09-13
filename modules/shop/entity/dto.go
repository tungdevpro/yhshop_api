package entity

import "coffee_api/commons"

type CreateShopDTO struct {
	Name  string         `json:"name" form:"name"`
	Icon  *commons.Image `json:"icon"`
	Image *commons.Image `json:"image"`
}
