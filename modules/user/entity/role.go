package entity

import "database/sql/driver"

type RoleAllowed string

const (
	admin  RoleAllowed = "admin"
	Seller RoleAllowed = "seller"
	Rider  RoleAllowed = "rider"
	Member RoleAllowed = "member"
)

func (st *RoleAllowed) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		*st = RoleAllowed(b)
	}
	return nil
}

func (st RoleAllowed) Value() (driver.Value, error) {
	return string(st), nil
}
