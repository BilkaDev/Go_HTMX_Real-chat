package store

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlStore struct {
	*gorm.DB
}

func NewSqliteStore() *SqlStore {

	db, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Conversation{}, &model.Message{})

	return &SqlStore{
		db,
	}
}
