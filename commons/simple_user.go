package commons

type SimpleUser struct {
	SQLModel `json:",inline"`
	Avatar   *Image `json:"avatar" gorm:"column:avatar;"`
	FullName string `json:"fullname" gorm:"column:fullname;"`

	// Status   StatusAllowed `json:"status" gorm:"column:status;type:ENUM('active','suspended','inactive');default:'active'"`
}

func (SimpleUser) TableName() string { return "users" }
