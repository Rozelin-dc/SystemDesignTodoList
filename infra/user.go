package infra

import (
	"database/sql"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userInfra struct {
	db *sqlx.DB
}

func NewUserInfra(db *sqlx.DB) repository.UserRepository {
	return &userInfra{db: db}
}

func (ui *userInfra) CreateUser(user *model.UserSimple) (*model.UserWithoutPass, error) {
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

	_, err = ui.db.Exec(
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

func (ui *userInfra) GetUser(userId string) (*model.UserWithoutPass, error) {
	user := model.UserWithoutPass{}
	err := ui.db.Get(
		&user,
		"SELECT `user_id`, `user_name` FROM `users` WHERE `user_id` = ?",
		userId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (ui *userInfra) EditUser(userId string, user *model.UserUpdate) (*model.UserWithoutPass, error) {
	updatedUser := model.User{}
	err := ui.db.Get(
		&updatedUser,
		"SELECT * FROM `users` WHERE `user_id` = ?",
		userId,
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

	_, err = ui.db.Exec(
		"UPDATE `users` SET `user_name` = ?, `password` = ? WHERE `user_id` = ?",
		user.UserName,
		string(hashed),
		userId,
	)
	if err != nil {
		return nil, err
	}

	return &model.UserWithoutPass{
		UserId:   userId,
		UserName: user.UserName,
	}, nil
}

func (ui *userInfra) DeleteUser(userId string, pass string) error {
	deletedUser := model.User{}
	err := ui.db.Get(
		&deletedUser,
		"SELECT * FROM `users` WHERE `user_id` = ?",
		userId,
	)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(deletedUser.Password), []byte(pass))
	if err != nil {
		return err
	}

	_, err = ui.db.Exec(
		"DELETE FROM `users` WHERE `user_id` = ?",
		userId,
	)

	return err
}

func (ui *userInfra) CheckRightUser(user *model.UserSimple) (*model.UserWithoutPass, error) {
	u := model.User{}
	err := ui.db.Get(
		&u,
		"SELECT * FROM `users` WHERE `user_id` = ?",
		user.UserName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return &model.UserWithoutPass{
		UserId:   u.UserId,
		UserName: u.UserName,
	}, nil
}

func (ui *userInfra) CheckUsedUserName(userName string) (*model.UserWithoutPass, error) {
	u := model.User{}
	err := ui.db.Get(
		&u,
		"SELECT * FROM `users` WHERE `user_name` = ?",
		userName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &model.UserWithoutPass{
		UserId:   u.UserId,
		UserName: u.UserName,
	}, nil
}
