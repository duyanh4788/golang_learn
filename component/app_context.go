package component

import (
	"golang_01/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnect() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, upProvider: upProvider}
}

func (ctx *appContext) GetMainDBConnect() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
