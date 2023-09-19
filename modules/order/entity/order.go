package entity

import "coffee_api/commons"

type Order struct {
	commons.SQLModel `json:",inline"`
}
