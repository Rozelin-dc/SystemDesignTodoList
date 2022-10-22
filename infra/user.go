package infra

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserInfra struct {
	DB *sqlx.DB
}

func NewUserInfra(db *sqlx.DB) *UserInfra {
	return &UserInfra{DB: db}
}

func (ui *UserInfra) CreateUser(user *model.UserSimple) (*model.UserWithoutPass, error) {
	pass := []byte(user.Password)
	hashed, err := bcrypt.GenerateFromPassword(pass, 10)
	if err != nil {
		return nil, err
	}

	uu, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	uuidStr := uu.String()

	_, err = ui.DB.Exec(
		"INSERT INTO `users` (`user_id`, `user_name`, `password`) VALUES (?, ?, ?)",
		uuidStr,
		user.UserName,
		string(hashed),
	)
	if err != nil {
		return nil, err
	}

	return &model.UserWithoutPass{
		UserId:   uuidStr,
		UserName: user.UserName,
	}, nil
}

func (ui *UserInfra) EditUser(user *model.UserUpdate) (*model.UserWithoutPass, error) {
	updatedUser := model.User{}
	err := ui.DB.Get(
		&updatedUser,
		"SELECT * FROM `users` WHERE `user_id` = ?",
		user.UserId,
	)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.NewPassword), 10)
	if err != nil {
		return nil, err
	}

	_, err = ui.DB.Exec(
		"UPDATE `users` SET `user_name` = ?, `password` = ? WHERE `user_id` = ?",
		user.UserName,
		string(hashed),
		user.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &model.UserWithoutPass{
		UserId:   user.UserId,
		UserName: user.UserName,
	}, nil
}

func (ui *UserInfra) DeleteUser(user *model.UserDelete) error {
	deletedUser := model.User{}
	err := ui.DB.Get(
		&deletedUser,
		"SELECT * FROM `users` WHERE `user_id` = ?",
		user.UserId,
	)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(deletedUser.Password), []byte(user.Password))
	if err != nil {
		return err
	}

	_, err = ui.DB.Exec(
		"DELETE FROM `users` WHERE `user_id` = ?",
		user.UserId,
	)

	return err
}
