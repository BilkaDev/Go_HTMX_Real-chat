package store

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
)

type MessageStore struct {
	Db *SqlStore
}

func NewMessageStore(db *SqlStore) *MessageStore {
	return &MessageStore{
		Db: db,
	}
}

func (ms *MessageStore) Create(m *model.Message) (*model.Message, error) {
	qm := ms.Db.Create(m)
	return m, qm.Error
}

func (ms *MessageStore) FindAll(conversationId uint) (*[]model.Message, error) {
	messages := []model.Message{}
	qm := ms.Db.Find(&messages, "conversation_id = ?", conversationId).Order("created_at DESC")
	return &messages, qm.Error

}
