package component

import (
	"golang_01/component/uploadprovider"
	"golang_01/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnect() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secret     string
	pubsub     pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secret string, pubsub pubsub.Pubsub) *appContext {
	return &appContext{db: db, upProvider: upProvider, secret: secret, pubsub: pubsub}
}

func (ctx *appContext) GetMainDBConnect() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appContext) SecretKey() string {
	return ctx.secret
}

func (ctx *appContext) GetPubSub() pubsub.Pubsub {
	return ctx.pubsub
}
