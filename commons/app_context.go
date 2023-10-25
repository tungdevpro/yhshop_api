package commons

import (
	"sync"

	"coffee_api/commons/mail"
	"coffee_api/configs"
	"coffee_api/pubsub"

	"gorm.io/gorm"
)

type AppContext struct {
	db     *gorm.DB
	L      *sync.RWMutex
	Cfg    *configs.Configuration
	pb     pubsub.Pubsub
	Mailer mail.EmailSender
}

func NewAppContext(db *gorm.DB, cfg *configs.Configuration, pb pubsub.Pubsub) *AppContext {
	mailer := mail.NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	return &AppContext{
		db:     db,
		L:      new(sync.RWMutex),
		Cfg:    cfg,
		pb:     pb,
		Mailer: mailer,
	}
}

func (appCtx *AppContext) GetDB() *gorm.DB { return appCtx.db.Session(&gorm.Session{}) }

func (appCtx *AppContext) GetPubSub() pubsub.Pubsub { return appCtx.pb }
