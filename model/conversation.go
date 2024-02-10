package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Participants []User    `gorm:"many2many:user_conversations;"`
	Messages     []Message `gorm:"foreignKey:ConversationID;"`
}
