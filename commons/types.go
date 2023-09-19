package commons

import "database/sql/driver"

type RoleAllowed string

const (
	admin   RoleAllowed = "admin"
	Seller  RoleAllowed = "seller"
	Shipper RoleAllowed = "shipper"
	Member  RoleAllowed = "member"
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

type StatusAllowed string

const (
	Suspended StatusAllowed = "suspended"
	Active    StatusAllowed = "active"
	Inactive  StatusAllowed = "inactive"
)

func (r *StatusAllowed) Scan(value interface{}) error {
	*r = StatusAllowed(value.([]byte))
	return nil
}

func (r StatusAllowed) Value() (driver.Value, error) {
	return string(r), nil
}
