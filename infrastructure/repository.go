package domain

import "github.com/Ad3bay0c/ddd-structure/domain/entity"

type DBInterface interface {
	CreateUser(user *entity.User) error
}
