package repository

import "github.com/Ad3bay0c/ddd-structure/domain/entity"

func (db *DB) CreateUser(user *entity.User) error {
	return db.Db.Error
}
