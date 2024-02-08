package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;type:varchar(100);default:null" form:"email"`
	FullName string `gorm:"required; not null;type:varchar(100);default:null" form:"fullName"`
	UserName string `gorm:"required; not null;type:varchar(100);default:null" form:"userName"`
	Password string `gorm:"required; not null;type:varchar(255);default:null" form:"password"`
	Avatar   string `gorm:"type:varchar(255);default:null" form:"avatar"`
	Gender   string `gorm:"required; not null;type:varchar(10);default:null" form:"gender"`
}
