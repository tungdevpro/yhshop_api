package commons

import (
	"time"
)

type SQLModel struct {
	Id        int        `json:"id,omitempty" gorm:"column:id;primary_key;"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
