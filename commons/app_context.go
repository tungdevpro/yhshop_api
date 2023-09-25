package commons

import (
	"coffee_api/configs"
	"coffee_api/pubsub"
	"sync"

	"gorm.io/gorm"
)

type AppContext struct {
	db  *gorm.DB
	L   *sync.RWMutex
	Cfg *configs.Configuration
	pb  pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, cfg *configs.Configuration, pb pubsub.Pubsub) *AppContext {
	return &AppContext{
		db:  db,
		L:   new(sync.RWMutex),
		Cfg: cfg,
		pb:  pb,
	}
}

func (ctx *AppContext) GetDB() *gorm.DB {
	return ctx.db.Session(&gorm.Session{})
}

func (app *AppContext) GetPubSub() pubsub.Pubsub {
	return app.pb
}
