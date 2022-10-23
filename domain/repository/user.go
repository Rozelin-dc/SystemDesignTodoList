package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type UserRepository interface {
	CreateUser(user *model.UserSimple) (*model.UserWithoutPass, error)
	GetUser(userId string) (*model.UserWithoutPass, error)
	EditUser(userId string, user *model.UserUpdate) (*model.UserWithoutPass, error)
	DeleteUser(userId string, pass string) error

	/** 正しいユーザー名とパスワードの組み合わせかどうか */
	CheckRightUser(user *model.UserSimple) (*model.UserWithoutPass, error)

	/** 既に使われているアカウント名かどうか */
	CheckUsedUserName(userName string) (*model.UserWithoutPass, error)
}
