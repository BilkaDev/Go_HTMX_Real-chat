package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ConversationID uint   `gorm:"required;"`
	SenderID       uint   `gorm:"required;"`
	ReceiverID     uint   `gorm:"required;" param:"id"`
	Message        string `gorm:"required;type:varchar(500);" form:"message" validate:"required,max=500"`
}
