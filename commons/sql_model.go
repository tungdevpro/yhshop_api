package commons

import (
	"time"

	"github.com/indrasaputra/hashids"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;primary_key;"`
	Uid       hashids.ID `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// func (sql *SQLModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	sql.Uid = hashids.ID(sql.Id)

// 	fmt.Println("uuid.....", sql.Uid)
// 	return nil
// }

func (sql *SQLModel) GenerateID() {
	sql.Uid = hashids.ID(sql.Id)
}
