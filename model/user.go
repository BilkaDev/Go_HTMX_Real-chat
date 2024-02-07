package model

import (
	"fmt"

	"github.com/bilkadev/Go_HTMX_Real-chat/types"
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

func (u *User) Create(store *types.SqlStore) (*User, error) {
	cu := store.Db.Create(u)
	if cu.Error != nil {
		fmt.Println("Create user failed.", cu.Error)
	} else {
		fmt.Println("create user: ")
	}
	return u, cu.Error
}

func (*User) FindAll(store *types.SqlStore) (*[]User, error) {
	var users []User
	query := store.Db.Find(&users)
	return &users, query.Error
}

func (u *User) FindOne(store *types.SqlStore, id uint) (*User, error) {
	query := store.Db.First(&u, id)
	return u, query.Error
}

func (u *User) FindOneByEmail(store *types.SqlStore, email string) (User, error) {
	var user User
	query := store.Db.First(&user, "email = ?", email)
	return user, query.Error
}
