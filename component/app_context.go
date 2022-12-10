package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnect() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (ctx *appContext) GetMainDBConnect() *gorm.DB {
	return ctx.db
}
