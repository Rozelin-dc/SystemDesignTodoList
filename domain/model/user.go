package model

// ログインとユーザー作成で使う
type UserSimple struct {
	UserName string `json:"userName" db:"user_name"`
	Password string `json:"password"`
}

type UserUpdate struct {
	UserName    string `json:"userName" db:"user_name"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type User struct {
	UserId   string `json:"userId" db:"user_id"`
	UserName string `json:"userName" db:"user_name"`
	Password string `json:"password"`
}

type UserWithoutPass struct {
	UserId   string `json:"userId" db:"user_id"`
	UserName string `json:"userName" db:"user_name"`
}
