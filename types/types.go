package types

import "gorm.io/gorm"

type SqlStore struct {
	Db *gorm.DB
}
