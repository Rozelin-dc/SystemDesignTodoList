package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type UserRepository interface {
	CreateUser(user *model.UserSimple) (*model.UserWithoutPass, error)
	GetUser(userId string) (*model.UserWithoutPass, error)
	EditUser(user *model.UserUpdate) (*model.UserWithoutPass, error)
	DeleteUser(user *model.UserDelete) error

	/** 正しいユーザー名とパスワードの組み合わせかどうか */
	CheckRightUser(user *model.UserSimple) (*model.UserWithoutPass, error)
}
