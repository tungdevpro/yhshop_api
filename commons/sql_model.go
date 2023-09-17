package commons

import (
	"fmt"
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;primary_key;"`
	Uid       *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (sql *SQLModel) GenUID() {
	uid := UID{
		LocalId: string(fmt.Sprintf("%d", sql.Id)),
	}

	x := uid.GenerateID()

	sql.Uid = &x
}
