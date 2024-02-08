package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;type:varchar(255);default:null" form:"email" validate:"email,required,min=3,max=255"`
	FullName string `gorm:"required; not null;type:varchar(100);default:null" form:"fullName" validate:"required,min=3,max=100"`
	UserName string `gorm:"required; not null;type:varchar(100);default:null" form:"userName" validate:"required,min=3,max=100"`
	Password string `gorm:"required; not null;type:varchar(255);default:null" form:"password" validate:"required,min=6,max=100"`
	Avatar   string `gorm:"type:varchar(255);default:null" form:"avatar"`
	Gender   string `gorm:"required; not null;type:varchar(10);default:null" form:"gender" validate:"required,oneof=male female" `
}
