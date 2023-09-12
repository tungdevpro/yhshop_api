package entity

import "coffee_api/commons"

type Shipper struct {
	*commons.SQLModel `json:",inline"`
}
