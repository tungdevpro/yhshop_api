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

func (appCtx *AppContext) GetDB() *gorm.DB {
	return appCtx.db.Session(&gorm.Session{})
}

func (appCtx *AppContext) GetPubSub() pubsub.Pubsub {
	return appCtx.pb
}
