package commons

import (
	"time"

	"github.com/indrasaputra/hashids"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	Uid       hashids.ID `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (sql *SQLModel) GenerateID() {
	sql.Uid = hashids.ID(sql.Id)
}
