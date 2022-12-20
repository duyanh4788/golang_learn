package component

import (
	"golang_01/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnect() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secret     string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secret string) *appContext {
	return &appContext{db: db, upProvider: upProvider, secret: secret}
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
