package commons

type SimpleUser struct {
	SQLModel `json:",inline"`
	Email    string        `json:"email"`
	Role     RoleAllowed   `json:"role"`
	Status   StatusAllowed `json:"status" gorm:"column:status;type:ENUM('active','suspended','inactive');default:'active'"`
}

func (SimpleUser) TableName() string { return "users" }

func (user *SimpleUser) IsActive() bool {
	return user.Status == Active
}
