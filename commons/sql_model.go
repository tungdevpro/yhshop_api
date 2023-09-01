package commons

import (
		"time"
)

type SQLModel struct {
	Id int64 `json:"id,omitemty"`
	CreatedAt *time.Time         `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
}
