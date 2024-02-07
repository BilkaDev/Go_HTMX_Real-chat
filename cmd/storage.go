package main

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteStore() *types.SqlStore {
	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	return &types.SqlStore{
		Db: db,
	}
}
