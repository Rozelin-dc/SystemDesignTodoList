package model

// ログインとユーザー作成で使う
type UserSimple struct {
	UserName string `json:"userName" db:"user_name" validate:"userName"`
	Password string `json:"password" validate:"password"`
}

type UserUpdate struct {
	UserName    string `json:"userName" db:"user_name" validate:"userName"`
	Password    string `json:"password" validate:"password"`
	NewPassword string `json:"newPassword" validate:"password"`
}

type User struct {
	UserId   string `json:"userId" db:"user_id"`
	UserName string `json:"userName" db:"user_name" validate:"userName"`
	Password string `json:"password" validate:"password"`
}

type UserWithoutPass struct {
	UserId   string `json:"userId" db:"user_id"`
	UserName string `json:"userName" db:"user_name" validate:"userName"`
}
