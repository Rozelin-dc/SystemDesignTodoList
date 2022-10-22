package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type UserRepository interface {
	CreateUser(user *model.UserSimple) (*model.UserWithoutPass, error)
	EditUser(user *model.UserUpdate) (*model.UserWithoutPass, error)
	DeleteUser(user *model.User) error
}
