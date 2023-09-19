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

type SimpleUser struct {
	SQLModel `json:",inline"`
	Email    string      `json:"email"`
	Role     RoleAllowed `json:"role"`
}

func (SimpleUser) TableName() string { return "users" }
