package commons

import (
	"sync"

	"coffee_api/commons/mail"
	"coffee_api/configs"
	"coffee_api/pubsub"
	"coffee_api/pubsub/pblocal"

	"gorm.io/gorm"
)

type AppContext struct {
	db     *gorm.DB
	L      *sync.RWMutex
	Cfg    *configs.Configuration
	ps     pubsub.Pubsub
	Mailer mail.EmailSender
}

func NewAppContext(db *gorm.DB, cfg *configs.Configuration) *AppContext {
	mailer := mail.NewGmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	return &AppContext{
		db:     db,
		L:      new(sync.RWMutex),
		Cfg:    cfg,
		ps:     pblocal.NewPubSub(),
		Mailer: mailer,
	}
}

func (appCtx *AppContext) GetDB() *gorm.DB { return appCtx.db.Session(&gorm.Session{}) }

func (ctx *AppContext) GetPubsub() pubsub.Pubsub {
	return ctx.ps
}
