package store

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
)

type ConversationStore struct {
	Db *SqlStore
}

func NewConversationStore(db *SqlStore) *ConversationStore {
	return &ConversationStore{
		Db: db,
	}
}

func (cs *ConversationStore) Create(c *model.Conversation, participants []uint) (*model.Conversation, error) {
	u1 := BuildFakeUser(participants[0])
	u2 := BuildFakeUser(participants[1])
	c.Participants = []model.User{u1, u2}
	qc := cs.Db.Create(c)

	return c, qc.Error
}

func (cs *ConversationStore) FindOne(c *model.Conversation, users []uint) (*model.Conversation, error) {
	qc := cs.Db.Table("conversations").
		Joins("JOIN user_conversations ON user_conversations.conversation_id = conversations.id").
		Where("user_conversations.user_id IN (?,?)", users[0], users[1]).
		Group("conversation_id").
		Having("COUNT(DISTINCT user_id) = 2").
		First(&c)
	return c, qc.Error
}

func (cs *ConversationStore) FindOneOrCreate(c *model.Conversation, message model.Message) (*model.Conversation, error) {
	qc, err := cs.FindOne(c, []uint{message.SenderID, message.ReceiverID})
	if err != nil {
		cs.Create(c, []uint{message.SenderID, message.ReceiverID})
	}
	return qc, err
}

func (cs *ConversationStore) Update(c *model.Conversation) (*model.Conversation, error) {
	qc := cs.Db.Save(&c)
	return c, qc.Error
}

func BuildFakeUser(id uint) model.User {
	u := model.User{
		Email:    "1",
		FullName: "1",
		Avatar:   "1",
		Gender:   "1",
	}
	u.UserName = "1"
	u.Password = "1"
	u.ID = id
	return u
}
