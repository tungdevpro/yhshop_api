package entity

import "coffee_api/commons"

type Rider struct {
	*commons.SQLModel `json:",inline"`
}
