package store

import (
	"fmt"

	"github.com/bilkadev/Go_HTMX_Real-chat/model"
)

type UserStore struct {
	Db *SqlStore
}

func NewUserStore(db *SqlStore) *UserStore {
	return &UserStore{
		Db: db,
	}
}

func (us *UserStore) Create(u *model.User) (*model.User, error) {
	cu := us.Db.Create(u)
	if cu.Error != nil {
		fmt.Println("Create user failed.", cu.Error)
	} else {
		fmt.Println("create user: ")
	}
	return u, cu.Error
}

func (us *UserStore) FindAll() (*[]model.User, error) {
	var users []model.User
	query := us.Db.Find(&users)
	return &users, query.Error
}

func (us *UserStore) FindOne(id uint) (*model.User, error) {
	var u model.User
	query := us.Db.First(&u, id)
	return &u, query.Error
}

func (us *UserStore) FindOneByEmail(email string) (*model.User, error) {
	var u model.User
	query := us.Db.First(&u, "email = ?", email)
	return &u, query.Error
}

func (us *UserStore) FindOneByUserName(userName string) (*model.User, error) {
	var u model.User
	query := us.Db.First(&u, "user_name = ?", userName)
	return &u, query.Error
}
