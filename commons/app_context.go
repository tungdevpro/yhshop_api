package commons

import (
	"coffee_api/configs"
	"sync"

	"gorm.io/gorm"
)

type AppContext struct {
	db  *gorm.DB
	L   *sync.RWMutex
	Cfg *configs.Configuration
	// upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, cfg *configs.Configuration) *AppContext {
	return &AppContext{
		db:  db,
		L:   new(sync.RWMutex),
		Cfg: cfg,
		// upProvider: upProvider,
	}
}

func (ctx *AppContext) GetDB() *gorm.DB {
	return ctx.db.Session(&gorm.Session{})
}
