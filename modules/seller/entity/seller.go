package entity

import "coffee_api/commons"

type Seller struct {
	*commons.SQLModel `json:",inline"`
}
