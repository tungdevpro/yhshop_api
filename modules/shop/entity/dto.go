package entity

import (
	"coffee_api/commons"
	"strings"
)

type CreateShopDTO struct {
	Name    string         `json:"name" form:"name"`
	OwnerId string         `json:"owner_id" form:"owner_id"`
	Icon    *commons.Image `json:"icon" form:"icon"`
	CityId  int            `json:"city_id" form:"city_id"`
}

func (c *CreateShopDTO) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	if len(c.Name) == 0 {
		return ErrNameIsEmpty
	}

	return nil
}
