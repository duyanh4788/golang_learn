package component

import (
	"golang_01/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnect() *gorm.DB
	UploadProvider() uploadprovider.UploadProvide
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvide
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvide) *appContext {
	return &appContext{db: db, upProvider: upProvider}
}

func (ctx *appContext) GetMainDBConnect() *gorm.DB {
	return ctx.db
}

func (ctx *appContext) UploadProvider() uploadprovider.UploadProvide {
	return ctx.upProvider
}
