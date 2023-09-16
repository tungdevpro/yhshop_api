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
		Id: sql.Id,
	}

	uid.Encrypt()

	sql.Uid = &uid
	e, _ := sql.Uid.MarshalJSON()

	fmt.Println("22....", e)
}
