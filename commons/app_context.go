package commons

import "gorm.io/gorm"

type AppContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *AppContext {
	return &AppContext{
		db: db,
	}
}

func (ctx *AppContext) GetDB() *gorm.DB {
	return ctx.db.Session(&gorm.Session{})
}
